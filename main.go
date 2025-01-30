package main

import (
	"embed"
	p2p_host "file-sharing/p2p"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	newHost, err := p2p_host.NewP2pHost()
	if err != nil {
		log.Println(err)
		return
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "file-sharing",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	err = newHost.Close()
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

}
