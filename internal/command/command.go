package command

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/Tauhoo/adon-desktop/internal/errors"
)

func RunWithDirectory(name string, args []string, directory string) (string, errors.Error) {
	command := exec.Command(name, args...)
	command.Dir = directory

	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr

	if err := command.Run(); err != nil {
		return "", errors.New(CommandRunFailCode, fmt.Sprint(err)+" : "+stderr.String())
	} else {
		return out.String(), nil
	}
}

func Run(name string, args []string) (string, errors.Error) {
	command := exec.Command(name, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr

	if err := command.Run(); err != nil {
		return "", errors.New(CommandRunFailCode, fmt.Sprint(err)+" : "+stderr.String())
	} else {
		return out.String(), nil
	}
}

func GetUserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
