package routes

import (
	"github.com/Tauhoo/adon-desktop/internal/messages"
	"github.com/Tauhoo/adon-desktop/internal/services"
	"github.com/asticode/go-astilectron"
)

var AddNewPlugin = func(service services.Service, m *astilectron.EventMessage) any {
	req, err := ReadEventMessage[services.PluginBuildInfo](m)
	if err != nil {
		return messages.NewResponseErrorMessage(err)
	}

	if err := service.AddNewPlugin(req.Data); err != nil {
		return messages.NewResponseErrorMessage(err)
	} else {
		return messages.NewResponseEmptyMessage()
	}
}

var GetPluginNameList = func(service services.Service, _ *astilectron.EventMessage) any {
	return messages.NewResponseMessage(service.GetPluginNameList())
}

var GetFunctionList = func(service services.Service, m *astilectron.EventMessage) any {
	req, err := ReadEventMessage[string](m)
	if err != nil {
		return messages.NewResponseErrorMessage(err)
	}

	if nameList, err := service.GetFunctionList(req.Data); err != nil {
		return messages.NewResponseErrorMessage(err)
	} else {
		return messages.NewResponseMessage(nameList)
	}
}

type GetFunctionReq struct {
	PluginName   string `json:"plugin_name"`
	FunctionName string `json:"function_name"`
}

var GetFunction = func(service services.Service, m *astilectron.EventMessage) any {
	req, err := ReadEventMessage[GetFunctionReq](m)
	if err != nil {
		return messages.NewResponseErrorMessage(err)
	}

	if function, err := service.GetFunction(req.Data.PluginName, req.Data.FunctionName); err != nil {
		return messages.NewResponseErrorMessage(err)
	} else {
		return messages.NewResponseMessage(function)
	}
}

var GetVariableList = func(service services.Service, m *astilectron.EventMessage) any {
	req, err := ReadEventMessage[string](m)
	if err != nil {
		return messages.NewResponseErrorMessage(err)
	}

	if nameList, err := service.GetVariableList(req.Data); err != nil {
		return messages.NewResponseErrorMessage(err)
	} else {
		return messages.NewResponseMessage(nameList)
	}
}

var GetAllGoBinPath = func(service services.Service, _ *astilectron.EventMessage) any {
	if nameList, err := service.GetAllGoBinPath(); err != nil {
		return messages.NewResponseErrorMessage(err)
	} else {
		return messages.NewResponseMessage(nameList)
	}
}
