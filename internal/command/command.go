package command

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/logs"
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

var sourceFiles = []string{
	".profile",
	".bash_login",
	".bash_profile",
	".bashrc",
	".zshrc",
}

func getSourceCommand() string {
	home := GetUserHomeDir()
	command := ""
	for _, file := range sourceFiles {
		bash := fmt.Sprintf("source %s/%s > /dev/null; ", home, file)
		command += bash
	}
	command += "echo $PATH"
	return command
}

func GetRealPath() (string, errors.Error) {
	if runtime.GOOS == "darwin" {
		bash := getSourceCommand()
		result, err := Run("bash", []string{"-c", bash})
		if err != nil {
			logs.InfoLogger.Printf("find path fail - error: %#v\n", err)
			return "", err
		}
		return strings.Trim(result, " \n"), nil
	}
	return os.Getenv("PATH"), nil
}
