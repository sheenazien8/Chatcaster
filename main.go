package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

var assets embed.FS

func main() {
	app := NewApp()

	// AppMenu := menu.NewMenu()

	err := wails.Run(&options.App{
		Title:           "Chatcaster",
		AlwaysOnTop:     true,
		Width:           400,
		Height:          400,
		Frameless:       true,
		CSSDragProperty: "widows",
		CSSDragValue:    "1",
		Windows: &windows.Options{
			WebviewIsTransparent: true,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{A: 0},
		OnStartup:        app.startup,
		// Menu:             AppMenu,
		Bind: []interface{}{
			app,
		},
		// WindowStartState: options.Maximised,
		Mac: &mac.Options{
			WebviewIsTransparent: true,
		},
		Linux: &linux.Options{},
		Debug: options.Debug{
			// OpenInspectorOnStartup: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
