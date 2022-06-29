package services

import (
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/plugin"
)

type PluginBuildInfo struct {
	ProjectPath string `json:"project_path"`
	GoPath      string `json:"go_path"`
	PluginName  string `json:"plugin_name"`
}

func (s service) AddNewPlugin(pluginBuildInfo PluginBuildInfo) errors.Error {
	info := plugin.BuildInfo{
		ProjectPath: pluginBuildInfo.ProjectPath,
		TargetPath:  s.config.WorkSpaceDirectory,
		GoPath:      pluginBuildInfo.GoPath,
		PluginName:  pluginBuildInfo.PluginName,
	}

	_, err := plugin.Build(info)
	if err != nil {
		return errors.New(BuildPluginFailCode, err)
	}

	return nil
}
