package main

import (
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
	defer astilectronInstance.Close()

	// Start astilectron
	astilectronInstance.Start()

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

	job := adon.NewJob()
	pluginManager := adon.NewPluginManager(job)
	service := services.New(pluginManager, window)

	routes.Regist(service, window)

	// Blocking pattern
	astilectronInstance.Wait()
}
