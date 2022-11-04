package main

import (
	"context"
	"embed"

	"github.com/dezhishen/windows-user-files-manager/pkg/application"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS
var app = application.NewApp()
var fileApp = application.NewFileApp()

func start(ctx context.Context) {
	app.Startup(ctx)
	fileApp.Startup(ctx)
}

func main() {
	// Create an instance of the app structure
	// Create application with options
	err := wails.Run(&options.App{
		Title:            "windows-user-files-manager",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        start,
		Bind: []interface{}{
			app, fileApp,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
