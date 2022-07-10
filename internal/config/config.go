package config

import (
	"net/url"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Env = string

const (
	DevEnv  = "dev"
	ProdEnv = "prod"
)

var env = DevEnv

type Config struct {
	Env                Env    `yaml:"env"`
	ClientLocation     string `yaml:"client-location"`
	AppName            string `yaml:"app-name"`
	AppIconDefaultPath string `yaml:"app-icon-default-path"`
	AppIconDarwinPath  string `yaml:"app-icon-darwin-path"`
	BaseDirectoryPath  string `yaml:"base-directory-path"`
	VersionAstilectron string `yaml:"version-astilectron"`
	VersionElectron    string `yaml:"version-electron"`
	WorkSpaceDirectory string `yaml:"work-space-directory"`
}

func NewFromFile(filePath string) (Config, error) {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(raw, &config); err != nil {
		return Config{}, err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	if _, err := url.ParseRequestURI(config.ClientLocation); err != nil {
		config.ClientLocation = path.Join(currentDir, config.ClientLocation)
	}

	config.AppIconDarwinPath = path.Join(currentDir, config.AppIconDarwinPath)
	config.AppIconDefaultPath = path.Join(currentDir, config.AppIconDefaultPath)
	config.WorkSpaceDirectory = path.Join(currentDir, config.WorkSpaceDirectory)
	config.BaseDirectoryPath = path.Join(currentDir, config.BaseDirectoryPath)

	return config, nil

}

func New() (Config, error) {
	switch env {
	case ProdEnv:
		return NewFromFile("resources/prod.config.yml")
	case DevEnv:
		fallthrough
	default:
		return NewFromFile("resources/dev.config.yml")
	}
}
