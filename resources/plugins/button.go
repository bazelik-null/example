package main

import (
	"os"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/bazelik-null/BBQDeploy/src/plugin"
	"github.com/bazelik-null/BBQDeploy/src/pluginapi"
)

func Init() {
	plugin.Global.RegisterEntry("PageInstall", AddBtnSteam)
}

func AddBtnSteam(payload interface{}) {
	pkg, ok := payload.(*pluginapi.PageInstallPackage)
	if !ok {
		return
	}

	btnSteam := widget.NewButtonWithIcon("Steam", theme.ComputerIcon(), func() {
		if runtime.GOOS == "windows" {
			pkg.Path = filepath.Join("C:\\", "Program Files (x86)", "Steam", "steamapps", "common", "example")
		} else {
			home := os.Getenv("HOME")
			pkg.Path = filepath.Join(home, ".steam", "root", "steamapps", "common", "example")
		}
		pkg.LabelPath.SetText("Chosen path: " + pkg.Path)
	})

	pkg.Container.Add(btnSteam)
}
