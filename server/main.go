package main

import (
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName: "adon",
		// AppIconDefaultPath: "<your .png icon>",  // If path is relative, it must be relative to the data directory
		// AppIconDarwinPath:  "<your .icns icon>", // Same here
		BaseDirectoryPath:  "dependencies",
		VersionAstilectron: "0.33.0",
		VersionElectron:    "6.1.2",
	})
	defer a.Close()

	// Start astilectron
	a.Start()

	w, _ := a.NewWindow("http://localhost:3000", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})

	w.Create()
	defer w.Close()

	// Blocking pattern
	a.Wait()
}
