package gocli

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Tauhoo/adon-desktop/internal/command"
	"github.com/Tauhoo/adon-desktop/internal/errors"
	"github.com/Tauhoo/adon-desktop/internal/logs"
)

var goBinMatcher = regexp.MustCompile(`^go([1-9](.[0-9]{1,2}){0,2})?$`)

func GetGOVERSION(binPath string) (string, errors.Error) {
	if result, err := command.Run(binPath, []string{"env", "GOVERSION"}); err != nil {
		return "", err
	} else {
		return strings.Trim(result, " \n"), nil
	}
}

func GetGOPATH(binPath string) (string, errors.Error) {
	if result, err := command.Run(binPath, []string{"env", "GOPATH"}); err != nil {
		return "", err
	} else {
		return strings.Trim(result, " \n"), nil
	}
}

func GetRealPath() (string, errors.Error) {
	home := command.GetUserHomeDir()
	bash := "source %s/.profile; source %s/.bash_profile; source %s/.zshrc; echo $PATH"
	bash = fmt.Sprintf(bash, home, home, home)
	result, err := command.Run("bash", []string{"-c", bash})
	if err != nil {
		logs.InfoLogger.Printf("find path fail - error: %#v\n", err)
		return "", err
	}
	return strings.Trim(result, " \n"), nil
}

func GetAllGoBin() ([]string, errors.Error) {
	result, err := GetRealPath()
	if err != nil {
		return nil, err
	}

	if err := os.Setenv("PATH", result); err != nil {
		return nil, errors.New(SetPATHEnvFailCode, err.Error())
	}

	binaries := []string{}

	if binPath, err := exec.LookPath("go"); err != nil {
		logs.ErrorLogger.Printf("look for go fail - error: %#v", err)
	} else {
		binaries = append(binaries, binPath)
	}

	if binPath, err := exec.LookPath("go1.18.3"); err != nil {
		logs.ErrorLogger.Printf("look for go1.18.3 fail - error: %#v", err)
	} else {
		binaries = append(binaries, binPath)
	}

	return binaries, nil
}
