package services

import (
	"os"
	"path"
	"reflect"

	"github.com/Tauhoo/adon"
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/gocli"
	"github.com/Tauhoo/adon-desktop/internal/logs"
	"github.com/Tauhoo/adon-desktop/internal/plugin"
)

type PluginBuildInfo struct {
	ProjectPath string `json:"project_path"`
	GoPath      string `json:"go_path"`
	PluginName  string `json:"plugin_name"`
}

func (s service) AddNewPlugin(pluginBuildInfo PluginBuildInfo) errors.Error {
	logs.InfoLogger.Printf("start add new plugin - info: %#v\n", pluginBuildInfo)
	info := plugin.BuildInfo{
		ProjectPath: pluginBuildInfo.ProjectPath,
		TargetPath:  s.config.WorkSpaceDirectory,
		GoPath:      pluginBuildInfo.GoPath,
		PluginName:  pluginBuildInfo.PluginName,
		Prefix:      pluginBuildInfo.PluginName,
	}

	targetFilename, err := plugin.Build(info)
	if err != nil {
		logs.ErrorLogger.Printf("build plugin fail - error: %#v\n", err.Error())
		return errors.New(BuildPluginFailCode, err.Error())
	}

	if err := s.pluginManager.LoadPluginFromFile(targetFilename); err != nil {
		logs.ErrorLogger.Printf("load plugin fail - error: %#v\n", err.Error())
		if rerr := os.Remove(targetFilename); rerr != nil {
			logs.ErrorLogger.Printf("delete plugin fail - error: %#v\n", rerr.Error())
		}
		return errors.New(LoadPluginFailCode, err.Error())
	}

	if err := s.api.PluginAdded(targetFilename); err != nil {
		logs.ErrorLogger.Printf("send plugin added event fail - error: %#v\n", err)
	}

	return nil
}

func (s service) GetPluginNameList() []string {
	names := []string{}

	for _, record := range s.pluginManager.GetPluginStorage().GetList() {
		names = append(names, record.Name)
	}

	return names
}

type Function struct {
	Name        string   `json:"name"`
	ArgTypes    []string `json:"arg_types"`
	ReturnTypes []string `json:"return_types"`
}

func (s service) GetFunctionList(pluginName string) ([]Function, errors.Error) {
	record, ok := s.pluginManager.GetPluginStorage().Find(pluginName)
	if !ok {
		return nil, errors.NewWithoutData(PluginNotFoundCode)
	}

	functionList := []Function{}
	for _, functionRecord := range record.Value.GetExecutorStorage().GetList() {

		argTypes := []string{}
		returnTypes := []string{}

		for _, param := range functionRecord.Value.GetFunction().GetParamList() {
			argTypes = append(argTypes, param.String())
		}

		for _, returnValue := range functionRecord.Value.GetFunction().GetReturnList() {
			returnTypes = append(returnTypes, returnValue.String())
		}

		functionList = append(functionList, Function{
			Name:        functionRecord.Name,
			ArgTypes:    argTypes,
			ReturnTypes: returnTypes,
		})
	}

	return functionList, nil
}

func (s service) GetFunction(pluginName string, functionName string) (Function, errors.Error) {
	logs.InfoLogger.Printf("find function %s in plugin %s\n", functionName, pluginName)
	pluginRecord, ok := s.pluginManager.GetPluginStorage().Find(pluginName)
	if !ok {
		logs.ErrorLogger.Printf("not found plugin %s\n", pluginName)
		return Function{}, errors.NewWithoutData(PluginNotFoundCode)
	}

	functionRecord, ok := pluginRecord.Value.GetExecutorStorage().Find(functionName)
	if !ok {
		logs.ErrorLogger.Printf("find function %s in plugin %s\n", functionName, pluginName)
		return Function{}, errors.NewWithoutData(FunctionNotFoundCode)
	}

	argTypeList := []string{}
	for _, argKind := range functionRecord.Value.GetFunction().GetParamList() {
		argTypeList = append(argTypeList, argKind.String())
	}

	returnTypeList := []string{}
	for _, returnKind := range functionRecord.Value.GetFunction().GetReturnList() {
		returnTypeList = append(returnTypeList, returnKind.String())
	}

	return Function{
		Name:        functionRecord.Name,
		ArgTypes:    argTypeList,
		ReturnTypes: returnTypeList,
	}, nil
}

type Variable struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (s service) GetVariableList(pluginName string) ([]Variable, errors.Error) {
	record, ok := s.pluginManager.GetPluginStorage().Find(pluginName)
	if !ok {
		return nil, errors.NewWithoutData(PluginNotFoundCode)
	}

	variableList := []Variable{}
	for _, varRecord := range record.Value.GetVariableStorage().GetList() {
		variableList = append(variableList, Variable{
			Name: varRecord.Name,
			Type: varRecord.Value.GetValue().Kind().String(),
		})
	}

	return variableList, nil
}

func (s service) ExecuteFunction(pluginName, functionName string, args []interface{}) errors.Error {
	logs.InfoLogger.Printf("start execute function - pluginName: %s, functionName: %s\n", pluginName, functionName)
	pluginRecord, ok := s.pluginManager.GetPluginStorage().Find(pluginName)
	if !ok {
		logs.ErrorLogger.Printf("not found plugin %s\n", pluginName)
		return errors.NewWithoutData(PluginNotFoundCode)
	}

	functionRecord, ok := pluginRecord.Value.GetExecutorStorage().Find(functionName)
	if !ok {
		logs.ErrorLogger.Printf("find function %s in plugin %s\n", functionName, pluginName)
		return errors.NewWithoutData(FunctionNotFoundCode)
	}

	argVariables := []adon.Variable{}

	params := functionRecord.Value.GetFunction().GetParamList()
	functionType := functionRecord.Value.GetFunction().GetValue().Type()

	if len(params) != len(args) {
		return errors.NewWithoutData(FunctionArgsInvalidCode)
	}

	for index, _ := range params {
		reflectValue := reflect.ValueOf(args[index]).Convert(functionType.In(index))
		logs.InfoLogger.Printf("%s ", reflectValue.Kind().String())
		argVariables = append(argVariables, adon.NewVariable(reflectValue))
	}

	logs.InfoLogger.Printf("execute function - pluginName: %s, functionName: %s\n", pluginName, functionName)
	functionRecord.Value.Execute(argVariables...)
	return nil
}
func (s service) LoadAllPlugin() {
	if err := s.pluginManager.LoadPluginFromFolder(s.config.WorkSpaceDirectory); err != nil {
		logs.ErrorLogger.Println(err.Error())
		os.Exit(1)
		return
	}

	for _, pluginRecord := range s.pluginManager.GetPluginStorage().GetList() {
		for _, executorRecord := range pluginRecord.Value.GetExecutorStorage().GetList() {
			stateEventPublisher := executorRecord.Value.GetStateEventPubliser()
			fn := func(state adon.ExecuteState, info any) {
				s.api.ExecutionStateChange(pluginRecord.Name, executorRecord.Name, state, info)
			}
			stateEventPublisher.Listen(adon.ExecuteDone, func(state adon.ExecuteState, info any) {
				variables := info.([]adon.Variable)
				infos := []any{}
				for _, variable := range variables {
					infos = append(infos, variable.GetValue().Interface())
				}

				s.api.ExecutionStateChange(pluginRecord.Name, executorRecord.Name, state, infos)
			})
			stateEventPublisher.Listen(adon.ExecuteRunning, fn)
			stateEventPublisher.Listen(adon.ExecuteError, fn)
			stateEventPublisher.Listen(adon.ExecuteIdle, fn)
		}
	}
}

func (s service) GetAllGoBinPath() ([]string, errors.Error) {
	return gocli.GetAllGoBin()
}

func (s service) DeletePlugin(name string) {
	s.pluginManager.GetPluginStorage().Delete(name)

	if rerr := os.Remove(path.Join(s.config.WorkSpaceDirectory, name)); rerr != nil {
		logs.ErrorLogger.Printf("delete plugin fail - error: %#v\n", rerr.Error())
	}

	if err := s.api.PluginDeleted(name); err != nil {
		logs.ErrorLogger.Printf("send plugin deleted event fail - error: %#v\n", err)
	}
}
