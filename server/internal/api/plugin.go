package api

import (
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/messages"
)

func (a api) PluginAdded(pluginName string) errors.Error {
	return a.send(messages.NewRequestMessage("route/plugin-added", pluginName))
}
