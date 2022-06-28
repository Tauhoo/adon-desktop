package services

import (
	"github.com/Tauhoo/adon"
	"github.com/Tauhoo/adon-desktop/internal/messages"
	"github.com/asticode/go-astilectron"
)

type Service interface {
	GetPluginList(_ messages.RequestMessage[any]) messages.ResponseMessage[[]string]
}

type service struct {
	pluginManager adon.PluginManager
	window        *astilectron.Window
}

func New(pluginManager adon.PluginManager, window *astilectron.Window) Service {
	return service{
		pluginManager: pluginManager,
		window:        window,
	}
}
