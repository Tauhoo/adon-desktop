package services

import (
	"os"

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
		return Function{}, errors.NewWithoutData(PluginNotFoundCode)
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

func (s service) GetAllGoBinPath() ([]string, errors.Error) {
	return gocli.GetAllGoBin()
}
