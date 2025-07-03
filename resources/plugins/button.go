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
		var path string
		if runtime.GOOS == "windows" {
			path = filepath.Join("C:\\", "Program Files (x86)", "Steam", "steamapps", "common", "ENA Dream BBQ")
		} else {
			home := os.Getenv("HOME")
			path = filepath.Join(home, ".steam", "root", "steamapps", "common", "ENA Dream BBQ")
		}
		pkg.LabelPath.SetText("Выбранный путь: " + path)
	})

	pkg.Container.Add(btnSteam)
}
