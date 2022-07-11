package main

import (
	"log"
	"os"

	"github.com/Tauhoo/adon"
	"github.com/Tauhoo/adon-desktop/internal/config"
	"github.com/Tauhoo/adon-desktop/internal/logs"
	"github.com/Tauhoo/adon-desktop/internal/routes"
	"github.com/Tauhoo/adon-desktop/internal/services"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	env := os.Getenv("ADON_ENV")
	conf, err := config.New(env)
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

	if _, err := os.Stat(conf.WorkSpaceDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(conf.WorkSpaceDirectory, 0777); err != nil {
			logs.ErrorLogger.Panicln("create work space directory fail - error: ", err.Error())
			return
		}
	}

	job := adon.NewJob()
	job.Start()
	logs.InfoLogger.Printf("start adon job")
	defer job.Stop()

	logs.InfoLogger.Printf("init adon plugin manager")
	pluginManager := adon.NewPluginManager(job)

	logs.InfoLogger.Printf("load plugin")
	service := services.New(pluginManager, window, conf)
	service.LoadAllPlugin()

	logs.InfoLogger.Printf("regist route")
	routes.Regist(service, window)

	logs.InfoLogger.Printf("create window at %s", conf.ClientLocation)
	window.Create()
	defer window.Close()

	if conf.Env == config.DevEnv {
		logs.InfoLogger.Printf("open dev tools")
		window.OpenDevTools()
	}

	logs.InfoLogger.Printf("start success")
	astilectronInstance.Wait()
}
