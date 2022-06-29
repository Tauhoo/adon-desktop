package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

type BuildInfo struct {
	ProjectPath string
	TargetPath  string
	GoPath      string
	PluginName  string
}

func Build(b BuildInfo) ([]byte, error) {
	filename := path.Join(b.TargetPath, b.PluginName+".so")

	command := exec.Command(
		b.GoPath,
		"build",
		"-buildmode=plugin",
		"-o",
		filename,
		"*.go")

	command.Dir = b.ProjectPath

	if err := os.RemoveAll(b.ProjectPath); err != nil {
		return nil, fmt.Errorf("%w - %s", ErrDeleteProjectFolder, err.Error())
	}

	return command.Output()

}
