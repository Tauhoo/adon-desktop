package routes

var handlers = map[string]Handler{
	"service/add-new-plugin":       AddNewPlugin,
	"service/get-function-list":    GetFunctionList,
	"service/get-variable-list":    GetVariableList,
	"service/get-plugin-name-list": GetPluginNameList,
}
