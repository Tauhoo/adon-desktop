package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

var env = "dev"

type Config struct {
	ClientLocation     string `yaml:"client-location"`
	AppName            string `yaml:"app-name"`
	AppIconDefaultPath string `yaml:"app-icon-default-path"`
	AppIconDarwinPath  string `yaml:"app-icon-darwin-path"`
	BaseDirectoryPath  string `yaml:"base-directory-path"`
	VersionAstilectron string `yaml:"version-astilectron"`
	VersionElectron    string `yaml:"version-electron"`
}

func NewConfigFromFile(path string) (Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(raw, &config); err != nil {
		return Config{}, err
	}

	return config, nil

}

func NewConfig() (Config, error) {
	switch env {
	case "prod":
		return NewConfigFromFile("resources/prod.config.yml")
	case "dev":
		fallthrough
	default:
		return NewConfigFromFile("resources/dev.config.yml")
	}
}
