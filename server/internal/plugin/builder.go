package plugin

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/Tauhoo/adon-desktop/internal/logs"
)

type BuildInfo struct {
	ProjectPath string
	TargetPath  string
	GoPath      string
	PluginName  string
}

func Build(b BuildInfo) (string, error) {
	targetFilename := path.Join(b.TargetPath, b.PluginName+".so")
	logs.InfoLogger.Printf("start build plugin - name: %s, projectPath: %s, target: %s\n", b.PluginName, b.ProjectPath, targetFilename)

	mainFiles, err := filepath.Glob(filepath.Join(b.ProjectPath, "*.go"))
	if err != nil {
		return "", errors.New(fmt.Sprint(err))
	}
	args := append([]string{"build", "-buildmode=plugin", "-o", targetFilename}, mainFiles...)

	command := exec.Command(b.GoPath, args...)
	command.Dir = b.ProjectPath

	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr

	if err = command.Run(); err != nil {
		return "", errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}

	return targetFilename, nil

}
