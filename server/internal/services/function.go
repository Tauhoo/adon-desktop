package services

import (
	"reflect"

	"github.com/Tauhoo/adon"
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/logs"
)

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

func (s service) ExecuteFunction(pluginName, functionName string, args []interface{}) errors.Error {
	logs.InfoLogger.Printf("start execute function - pluginName: %s, functionName: %s\n", pluginName, functionName)
	pluginRecord, ok := s.pluginManager.GetPluginStorage().Find(pluginName)
	if !ok {
		logs.ErrorLogger.Printf("not found plugin %s\n", pluginName)
		return errors.NewWithoutData(PluginNotFoundCode)
	}

	functionRecord, ok := pluginRecord.Value.GetExecutorStorage().Find(functionName)
	if !ok {
		logs.ErrorLogger.Printf("find function %s in plugin %s\n", functionName, pluginRecord.Name)
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
		pointerReflectValue := reflect.New(reflectValue.Type())
		pointerReflectValue.Elem().Set(reflectValue)
		argVariables = append(argVariables, adon.NewVariableFromPointer(pointerReflectValue))
	}

	logs.InfoLogger.Printf("execute function - pluginName: %s, functionName: %s\n", pluginRecord.Name, functionRecord.Name)
	functionRecord.Value.Execute(argVariables...)
	return nil
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
