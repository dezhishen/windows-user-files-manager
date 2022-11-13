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
var fileApp = application.NewFileApp()

func start(ctx context.Context) {
	application.Startup(ctx)
}

func close(ctx context.Context) bool {
	return application.Close(ctx)
}

func shutdown(ctx context.Context) {
	application.Shutdown(ctx)
}

func getBinds() []interface{} {
	apps := application.GetApps()
	binds := make([]interface{}, len(apps))
	for i, app := range apps {
		binds[i] = app
	}
	return binds
}

func main() {
	// Create an instance of the app structure
	// Create application with options
	err := wails.Run(&options.App{
		Title:         "文件管理工具",
		Width:         1024,
		Height:        768,
		Assets:        assets,
		OnStartup:     start,
		OnShutdown:    shutdown,
		OnBeforeClose: close,
		Bind:          getBinds(),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
