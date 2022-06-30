package services

import (
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/logs"
	"github.com/Tauhoo/adon-desktop/internal/plugin"
)

type PluginBuildInfo struct {
	ProjectPath string `json:"project_path"`
	GoPath      string `json:"go_path"`
	PluginName  string `json:"plugin_name"`
}

func (s service) AddNewPlugin(pluginBuildInfo PluginBuildInfo) errors.Error {
	logs.InfoLogger.Printf("start add new plugin - info: %#v\n", pluginBuildInfo)
	info := plugin.BuildInfo{
		ProjectPath: pluginBuildInfo.ProjectPath,
		TargetPath:  s.config.WorkSpaceDirectory,
		GoPath:      pluginBuildInfo.GoPath,
		PluginName:  pluginBuildInfo.PluginName,
	}

	targetFilename, err := plugin.Build(info)
	if err != nil {
		logs.ErrorLogger.Printf("build plugin fail - error: %#v\n", err.Error())
		return errors.New(BuildPluginFailCode, err.Error())
	}

	if err := s.pluginManager.LoadPluginFromFile(targetFilename); err != nil {
		logs.ErrorLogger.Printf("build plugin fail - error: %#v\n", err.Error())
		return errors.New(LoadPluginFailCode, err.Error())
	}

	return nil
}
