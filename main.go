package main

import (
	"embed"
	p2p_host "file-sharing/p2p"
	"file-sharing/pkg/app"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	newHost, currentMultiAddress, err := p2p_host.NewP2pHost()

	// Create an instance of the app structure
	app := app.NewApp(currentMultiAddress)

	if err != nil {
		log.Println(err)
		return
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "FileSharing",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
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
