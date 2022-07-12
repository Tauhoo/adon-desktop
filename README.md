<img src="./resources/logo.png" width="100px" style="margin-bottom: 20px;"/>

# Adon
Adon is basic Go plugin interpretor. It convert Go plugin to GUI.

# Installation
## Go binary
To compile plugin, you need to install Go version `1.18.3` first. Because, Adon use `go` or `go1.18.3` command to compile plugins.
## Adon
Download here

# Plugin
To create a plugin, you can just init basic go project with go mod init.
```bash
mkdir plugin-folder
cd plugin-folder
go mod init example.com/m
touch main.go
```

### main.go
```go
package main

var PI float32 = 3.14

func CircleArea(radius float32) float32 {
	return PI*radius*radius
}

func CircleCircumference(radius float32) float32 {
	return 2*PI*radius
}
```

### go.mod
```
module example.com/m

go 1.18
```

Then open Adon and create a plugin with Create Plugin button. 

<img src="readme-assets/create-button.png" width="200px">

Fill plugin name, plugin project path and click build

<img src="readme-assets/create-modal.png" height="150px">

New plugin show at the left

<img src="readme-assets/plugin-menu.png" width="200px">

# Road map
In the future, Adon has plans to improve. The list below are the features.
- [ ] Go struct support.
- [ ] Create new plugin from git repo.
- [ ] More type of input such as option input.
- [ ] Variable saving system.
- [ ] Function searching