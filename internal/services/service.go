package services

import (
	"github.com/Tauhoo/adon"
	"github.com/Tauhoo/adon-desktop/internal/api"
	"github.com/Tauhoo/adon-desktop/internal/config"
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/asticode/go-astilectron"
)

type Service interface {
	AddNewPlugin(pluginBuildInfo PluginBuildInfo) errors.Error
	GetPluginNameList() []string
	GetFunctionList(pluginName string) ([]Function, errors.Error)
	GetVariableList(pluginName string) ([]Variable, errors.Error)
	SetVariable(pluginName string, variableMap map[string]interface{}) errors.Error
	GetFunction(pluginName string, functionName string) (Function, errors.Error)
	DeletePlugin(name string)
	ExecuteFunction(pluginName, functionName string, args []interface{}) errors.Error
	LoadAllPlugin()

	GetAllGoBinPath() ([]string, errors.Error)
}

type service struct {
	pluginManager adon.PluginManager
	window        *astilectron.Window
	config        config.Config
	api           api.API
}

func New(pluginManager adon.PluginManager, window *astilectron.Window, conf config.Config) Service {
	return service{
		pluginManager: pluginManager,
		window:        window,
		config:        conf,
		api:           api.New(window),
	}
}
