package gocli

import (
	"io/ioutil"
	"path"
	"regexp"
	"strings"

	"github.com/Tauhoo/adon-desktop/internal/command"
	"github.com/Tauhoo/adon-desktop/internal/errors"
)

var goBinMatcher = regexp.MustCompile(`^go([1-9](.[0-9]{1,2}){0,2})?$`)

func GetGOPATH() (string, errors.Error) {
	if result, err := command.Run("go", []string{"env", "GOPATH"}); err != nil {
		return "", err
	} else {
		return strings.Trim(result, " \n"), nil
	}
}

func GetAllGoBin() ([]string, errors.Error) {
	binaries := []string{}

	if result, err := command.Run("which", []string{"go"}); err != nil {
		return nil, err
	} else {
		binaries = append(binaries, strings.Trim(result, " \n"))
	}

	gopath, err := GetGOPATH()
	if err != nil {
		return nil, err
	}

	gopath = path.Join(gopath, "bin")
	files, rawerr := ioutil.ReadDir(gopath)
	if rawerr != nil {
		return nil, errors.New(ReadDirFailCode, rawerr.Error())
	}

	for _, file := range files {
		filename := file.Name()
		if goBinMatcher.MatchString(filename) {
			binaries = append(binaries, path.Join(gopath, filename))
		}
	}

	return binaries, nil
}
