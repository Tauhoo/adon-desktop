package config

import (
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/Tauhoo/adon-desktop/internal/logs"
	"gopkg.in/yaml.v3"
)

type Env = string

var (
	DevEnv  Env = "dev"
	ProdEnv Env = "prod"
)

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

func GetBinaryDirectory(env Env) (string, error) {
	if env == DevEnv {
		currentDir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return currentDir, nil
	} else {
		ex, err := os.Executable()
		if err != nil {
			return "", err
		}
		exPath := filepath.Dir(ex)
		return exPath, nil
	}
}

func NewFromFile(execDir, filePath string) (Config, error) {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(raw, &config); err != nil {
		return Config{}, err
	}

	logs.InfoLogger.Printf("current directory - %s\n", execDir)

	if _, err := url.ParseRequestURI(config.ClientLocation); err != nil {
		logs.ErrorLogger.Printf("parse url fail - %s\n", err.Error())
		config.ClientLocation = path.Join(execDir, config.ClientLocation)
	}

	config.AppIconDarwinPath = path.Join(execDir, config.AppIconDarwinPath)
	config.AppIconDefaultPath = path.Join(execDir, config.AppIconDefaultPath)
	config.WorkSpaceDirectory = path.Join(execDir, config.WorkSpaceDirectory)
	config.BaseDirectoryPath = path.Join(execDir, config.BaseDirectoryPath)
	logs.InfoLogger.Printf("config value - %#v\n", config)
	return config, nil

}

func New(env string) (Config, error) {
	currentDir, err := GetBinaryDirectory(env)
	if err != nil {
		return Config{}, err
	}

	logs.InfoLogger.Printf("config value env (%s, %s) - %s\n", ProdEnv, DevEnv, env)
	switch env {
	case DevEnv:
		return NewFromFile(currentDir, path.Join(currentDir, "resources/dev.config.yml"))
	case ProdEnv:
		fallthrough
	default:
		return NewFromFile(currentDir, path.Join(currentDir, "resources/prod.config.yml"))
	}
}
