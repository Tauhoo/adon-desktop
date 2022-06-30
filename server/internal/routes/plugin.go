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
