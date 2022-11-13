package application

import (
	"context"

	"github.com/dezhishen/windows-user-files-manager/pkg/config"
)

type ConfigApp struct {
	ctx context.Context
}

func NewConfigApp() *ConfigApp {
	return &ConfigApp{}
}

func init() {
	AddApp(NewConfigApp())
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *ConfigApp) Startup(ctx context.Context) {
	app.ctx = ctx
}

func (app *ConfigApp) Close(ctx context.Context) {
}

func (app *ConfigApp) Shutdown(ctx context.Context) {
	config.Unlock()
}

func (app *ConfigApp) GetConfigValue(key string) (string, error) {
	return config.GetConfigValue(key)
}

func (app *ConfigApp) PutConfigValue(key, value string) error {
	return config.PutConfigValue(key, value)
}
func (app *ConfigApp) PutAllValues(values map[string]string) error {
	return config.PutAllValues(values)
}

func (app *ConfigApp) GetAllValues() (map[string]string, error) {
	return config.GetConfigAllValues()
}
