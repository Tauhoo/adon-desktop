package services

import (
	"os"
	"path"

	"github.com/Tauhoo/adon"
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/gocli"
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
		Prefix:      pluginBuildInfo.PluginName,
	}

	targetFilename, err := plugin.Build(info)
	if err != nil {
		logs.ErrorLogger.Printf("build plugin fail - error: %#v\n", err.Error())
		return errors.New(BuildPluginFailCode, err.Error())
	}

	if err := s.pluginManager.LoadPluginFromFile(targetFilename); err != nil {
		logs.ErrorLogger.Printf("load plugin fail - error: %#v\n", err.Error())
		if rerr := os.Remove(targetFilename); rerr != nil {
			logs.ErrorLogger.Printf("delete plugin fail - error: %#v\n", rerr.Error())
		}
		return errors.New(LoadPluginFailCode, err.Error())
	}

	pluginName := path.Base(targetFilename)

	if err := s.api.PluginAdded(pluginName); err != nil {
		logs.ErrorLogger.Printf("send plugin added event fail - error: %#v\n", err)
	}

	pluginRecord, ok := s.pluginManager.GetPluginStorage().Find(pluginName)
	if !ok {
		return errors.NewWithoutData(PluginNotFoundCode)
	}

	for _, executorRecord := range pluginRecord.Value.GetExecutorStorage().GetList() {
		stateEventPublisher := executorRecord.Value.GetStateEventPubliser()
		pluginName := pluginRecord.Name
		executorName := executorRecord.Name
		fn := func(state adon.ExecuteState, info any) {
			s.api.ExecutionStateChange(pluginName, executorName, state, info)
		}
		stateEventPublisher.Listen(adon.ExecuteDone, func(state adon.ExecuteState, info any) {
			variables := info.([]adon.Variable)
			infos := []any{}
			for _, variable := range variables {
				infos = append(infos, variable.GetValue().Interface())
			}

			s.api.ExecutionStateChange(pluginName, executorName, state, infos)
		})
		stateEventPublisher.Listen(adon.ExecuteRunning, fn)
		stateEventPublisher.Listen(adon.ExecuteError, fn)
		stateEventPublisher.Listen(adon.ExecuteIdle, fn)
	}

	return nil
}

func (s service) GetPluginNameList() []string {
	names := []string{}

	for _, record := range s.pluginManager.GetPluginStorage().GetList() {
		names = append(names, record.Name)
	}

	return names
}

func (s service) LoadAllPlugin() {
	if err := s.pluginManager.LoadPluginFromFolder(s.config.WorkSpaceDirectory); err != nil {
		logs.ErrorLogger.Println(err.Error())
		os.Exit(1)
		return
	}

	for _, pluginRecord := range s.pluginManager.GetPluginStorage().GetList() {
		for _, executorRecord := range pluginRecord.Value.GetExecutorStorage().GetList() {
			stateEventPublisher := executorRecord.Value.GetStateEventPubliser()
			pluginName := pluginRecord.Name
			functionName := executorRecord.Name
			fn := func(state adon.ExecuteState, info any) {
				s.api.ExecutionStateChange(pluginName, functionName, state, info)
			}
			stateEventPublisher.Listen(adon.ExecuteDone, func(state adon.ExecuteState, info any) {
				variables := info.([]adon.Variable)
				infos := []any{}
				for _, variable := range variables {
					infos = append(infos, variable.GetValue().Interface())
				}
				s.api.ExecutionStateChange(pluginName, functionName, state, infos)
			})
			stateEventPublisher.Listen(adon.ExecuteRunning, fn)
			stateEventPublisher.Listen(adon.ExecuteError, fn)
			stateEventPublisher.Listen(adon.ExecuteIdle, fn)
		}
	}
}

func (s service) GetAllGoBinPath() ([]string, errors.Error) {
	return gocli.GetAllGoBin()
}

func (s service) DeletePlugin(name string) {
	s.pluginManager.GetPluginStorage().Delete(name)

	if rerr := os.Remove(path.Join(s.config.WorkSpaceDirectory, name)); rerr != nil {
		logs.ErrorLogger.Printf("delete plugin fail - error: %#v\n", rerr.Error())
	}

	if err := s.api.PluginDeleted(name); err != nil {
		logs.ErrorLogger.Printf("send plugin deleted event fail - error: %#v\n", err)
	}
}
