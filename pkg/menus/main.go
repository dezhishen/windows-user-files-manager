package menus

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var _menu = menu.NewMenu()
var _ctx context.Context

func Startup(ctx context.Context) {
	_ctx = ctx
}

func GetMenu() *menu.Menu {
	return _menu
}

func init() {
	InitMenu()
}

func InitMenu() {
	menu := _menu.AddSubmenu("菜单")
	menu.AddText("设置", nil, openSettings)
	menu.AddText("增加页签", nil, openNewRootpathdialog)

}

func openSettings(callbackData *menu.CallbackData) {
	runtime.EventsEmit(_ctx, "openSettingDialog")
}

func openNewRootpathdialog(callbackData *menu.CallbackData) {
	runtime.EventsEmit(_ctx, "openNewRootpathdialog")
}
