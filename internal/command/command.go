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

func GetRealEnv() (map[string]string, errors.Error) {
	args := []string{
		"-ilc",
		`echo -n "_SHELL_ENV_DELIMITER_"; env; echo -n "_SHELL_ENV_DELIMITER_"; exit`,
	}

	rawenv, err := Run(os.Getenv("SHELL"), args)
	if err != nil {
		logs.InfoLogger.Printf("find path fail - error: %#v\n", err)
		return nil, err
	}

	envEntries := strings.Split(rawenv, "\n")

	noneEmptyEnvEntries := []string{}

	for _, envEntry := range envEntries {
		if envEntry != "" {
			noneEmptyEnvEntries = append(noneEmptyEnvEntries, envEntry)
		}
	}

	result := map[string]string{}
	for _, noneEmptyEnvEntry := range noneEmptyEnvEntries {
		envKeyValue := strings.Split(noneEmptyEnvEntry, "=")
		if len(envKeyValue) < 2 {
			continue
		}
		valueList := []string{}
		for i := 1; i < len(envKeyValue); i++ {
			valueList = append(valueList, envKeyValue[i])
		}
		result[envKeyValue[0]] = strings.Join(valueList, "=")
	}

	return result, nil
}

func GetRealPath() (string, errors.Error) {
	if runtime.GOOS != "windows" {
		env, err := GetRealEnv()
		if err != nil {
			return "", err
		}
		path, ok := env["PATH"]
		if ok {
			return path, nil
		}
		return os.Getenv("PATH"), nil
	}
	return os.Getenv("PATH"), nil
}
