package main

import (
	"context"
	"embed"

	"github.com/dezhishen/windows-user-files-manager/pkg/application"
	"github.com/dezhishen/windows-user-files-manager/pkg/menus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func startup(ctx context.Context) {
	application.Startup(ctx)
	menus.Startup(ctx)
}

func beforeClose(ctx context.Context) bool {
	return application.BeforeClose(ctx)
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
		OnStartup:     startup,
		OnShutdown:    shutdown,
		OnBeforeClose: beforeClose,
		Bind:          getBinds(),
		Menu:          menus.GetMenu(),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
