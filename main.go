package main

import (
	"embed"

	"github.com/tidwall/gjson"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed all:wails.json
var wailsJSON string

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Get version from wails.json
	version := gjson.Get(wailsJSON, "info.productVersion").String()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "SIMAppletManager v" + version,
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
