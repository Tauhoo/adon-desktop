package plugin

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"

	"github.com/Tauhoo/adon-desktop/internal/command"
	"github.com/Tauhoo/adon-desktop/internal/logs"
)

type BuildInfo struct {
	ProjectPath string
	TargetPath  string
	GoPath      string
	PluginName  string
	Prefix      string
}

func Build(b BuildInfo) (string, error) {
	targetFilename := path.Join(b.TargetPath, b.PluginName+".so")
	logs.InfoLogger.Printf("start build plugin - name: %s, projectPath: %s, target: %s\n", b.PluginName, b.ProjectPath, targetFilename)

	if _, err := command.RunWithDirectory(b.GoPath, []string{"mod", "tidy"}, b.ProjectPath); err != nil {
		return "", errors.New(fmt.Sprintf("go mod tidy - error: %#v", err))
	}

	mainFiles, rawerr := filepath.Glob(filepath.Join(b.ProjectPath, "*.go"))
	if rawerr != nil {
		return "", errors.New(fmt.Sprint(rawerr))
	}

	args := append([]string{"build", "-buildmode=plugin", fmt.Sprintf("-ldflags=\"-pluginpath=plugin/%s\"", b.Prefix), "-o", targetFilename}, mainFiles...)
	if _, err := command.RunWithDirectory(b.GoPath, args, b.ProjectPath); err != nil {
		return "", errors.New(fmt.Sprintf("build fail - error: %#v", err))
	}

	return targetFilename, nil
}
