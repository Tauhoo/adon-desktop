package gocli

import (
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

func GetAllGoBin() ([]string, errors.Error) {
	result, err := command.GetRealPath()
	if err != nil {
		return nil, err
	}

	logs.InfoLogger.Printf("real path - error: %s", result)

	if err := os.Setenv("PATH", result); err != nil {
		return nil, errors.New(SetPATHEnvFailCode, err.Error())
	}

	binaries := []string{}
	commands := []string{"go", "go1.18.3"}

	for _, cmd := range commands {
		binPath, rawerr := exec.LookPath(cmd)
		if rawerr != nil {
			logs.ErrorLogger.Printf("look for %s fail - error: %#v", cmd, rawerr)
		}

		version, err := GetGOVERSION(binPath)
		if err != nil {
			logs.ErrorLogger.Printf("get go version from bin fail - binary: %s", cmd)
			continue
		}

		if version != "go1.18.3" {
			continue
		}

		binaries = append(binaries, binPath)
	}

	return binaries, nil
}
