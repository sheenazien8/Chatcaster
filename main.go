package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	runtime_walls "github.com/wailsapp/wails/v2/pkg/runtime"
)

var assets embed.FS

func main() {
	app := NewApp()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Settings", keys.CmdOrCtrl(","), func(cd *menu.CallbackData) {
		runtime_walls.EventsEmit(app.ctx, "navigatedTo", "settings")
	})
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime_walls.Quit(app.ctx)
	})

	if runtime.GOOS == "darwin" {
		AppMenu.Append(menu.EditMenu())
		AppMenu.Append(menu.WindowMenu())
	}

	err := wails.Run(&options.App{
		Title:       "Chatcaster",
		AlwaysOnTop: true,
		Width:       300,
		Height:      200,
		Frameless:   true,
		// CSSDragProperty: "widows",
		// CSSDragValue:    "1",
		Windows: &windows.Options{
			WebviewIsTransparent: true,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{A: 0},
		OnStartup:        app.startup,
		Menu:             AppMenu,
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
