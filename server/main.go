package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {

	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("DEBUG: %#v\n", config)
	astilectronInstance, err := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            config.AppName,
		AppIconDefaultPath: config.AppIconDefaultPath,
		AppIconDarwinPath:  config.AppIconDarwinPath,
		BaseDirectoryPath:  config.BaseDirectoryPath,
		VersionAstilectron: config.VersionAstilectron,
		VersionElectron:    config.VersionElectron,
	})
	defer astilectronInstance.Close()

	// Start astilectron
	astilectronInstance.Start()

	window, err := astilectronInstance.NewWindow(config.ClientLocation, &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})
	if err != nil {
		panic(err)
	}

	window.Create()
	defer window.Close()

	// Blocking pattern
	astilectronInstance.Wait()
}
