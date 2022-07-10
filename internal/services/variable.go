package services

import (
	"reflect"

	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/logs"
)

type Variable struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func (s service) GetVariableList(pluginName string) ([]Variable, errors.Error) {
	record, ok := s.pluginManager.GetPluginStorage().Find(pluginName)
	if !ok {
		return nil, errors.NewWithoutData(PluginNotFoundCode)
	}

	variableList := []Variable{}
	for _, varRecord := range record.Value.GetVariableStorage().GetList() {
		variableList = append(variableList, Variable{
			Name:  varRecord.Name,
			Type:  varRecord.Value.GetValue().Kind().String(),
			Value: varRecord.Value.GetValue().Interface(),
		})
	}

	return variableList, nil
}

func (s service) SetVariable(pluginName string, variableMap map[string]interface{}) errors.Error {
	pluginRecord, ok := s.pluginManager.GetPluginStorage().Find(pluginName)
	if !ok {
		logs.ErrorLogger.Printf("not found plugin %s\n", pluginName)
		return errors.NewWithoutData(PluginNotFoundCode)
	}

	for key, value := range variableMap {
		variableRecord, ok := pluginRecord.Value.GetVariableStorage().Find(key)
		if !ok {
			logs.ErrorLogger.Printf("not found variable %s in plugin %s\n", key, pluginName)
			continue
		}

		varType := variableRecord.Value.GetValue().Type()
		reflectVar := reflect.ValueOf(value)
		if reflectVar.CanConvert(varType) {
			variableRecord.Value.GetValue().Set(reflectVar.Convert(varType))
		} else {
			logs.ErrorLogger.Printf("cannot convert variable %s to %s kind in plugin %s\n", key, reflectVar.Kind().String(), pluginName)
		}
	}

	return nil
}
