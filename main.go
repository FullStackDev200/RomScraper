package main

import (
	"embed"
	"log"

	sc "github.com/FullStackDev200/RomScraper/scraping"
	uc "github.com/FullStackDev200/RomScraper/userconfig"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	configstore, err := uc.NewConfigStore()
	if err != nil {
		log.Println("Error opening configstore:", err)
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "RomScraper",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			configstore,
		},
		EnumBind: []interface{}{
			sc.AllPlatforms,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
