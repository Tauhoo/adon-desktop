package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tauhoo/adon"
	"github.com/Tauhoo/adon-desktop/internal/config"
	"github.com/Tauhoo/adon-desktop/internal/routes"
	"github.com/Tauhoo/adon-desktop/internal/services"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	astilectronInstance, err := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            conf.AppName,
		AppIconDefaultPath: conf.AppIconDefaultPath,
		AppIconDarwinPath:  conf.AppIconDarwinPath,
		BaseDirectoryPath:  conf.BaseDirectoryPath,
		VersionAstilectron: conf.VersionAstilectron,
		VersionElectron:    conf.VersionElectron,
	})

	astilectronInstance.Start()
	defer astilectronInstance.Close()
	// Start astilectron

	window, err := astilectronInstance.NewWindow(conf.ClientLocation, &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})
	if err != nil {
		panic(err)
	}

	window.Create()
	defer window.Close()

	if conf.Env == config.DevEnv {
		window.OpenDevTools()
	}

	if _, err := os.Stat(conf.WorkSpaceDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(conf.WorkSpaceDirectory, os.ModeDir.Perm()); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	job := adon.NewJob()
	pluginManager := adon.NewPluginManager(job)
	service := services.New(pluginManager, window, conf)

	routes.Regist(service, window)

	// Blocking pattern
	astilectronInstance.Wait()
}
