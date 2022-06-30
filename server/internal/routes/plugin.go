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
