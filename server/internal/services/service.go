package services

import (
	"github.com/Tauhoo/adon"
	"github.com/asticode/go-astilectron"
)

type Service interface{}

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
