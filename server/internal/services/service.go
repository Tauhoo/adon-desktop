package services

import (
	"github.com/Tauhoo/adon"
	"github.com/Tauhoo/adon-desktop/internal/config"
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/asticode/go-astilectron"
)

type Service interface {
	AddNewPlugin(pluginBuildInfo PluginBuildInfo) errors.Error
}

type service struct {
	pluginManager adon.PluginManager
	window        *astilectron.Window
	config        config.Config
}

func New(pluginManager adon.PluginManager, window *astilectron.Window, conf config.Config) Service {
	return service{
		pluginManager: pluginManager,
		window:        window,
		config:        conf,
	}
}
