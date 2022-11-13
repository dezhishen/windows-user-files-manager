package application

import (
	"context"
	"log"

	"github.com/dezhishen/windows-user-files-manager/pkg/config"
	"github.com/dezhishen/windows-user-files-manager/pkg/fileutil"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FileApp struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewFileApp() *FileApp {
	return &FileApp{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *FileApp) Startup(ctx context.Context) {
	app.ctx = ctx
}

func (app *FileApp) Close(ctx context.Context) {}

func (app *FileApp) Shutdown(ctx context.Context) {}

// Greet returns a greeting for the given name
func (app *FileApp) GetAllDir() []*fileutil.Dir {
	result, err := fileutil.GetFile(fileutil.GetUserDir())
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (app *FileApp) GetByRootpath(rootpath string) []*fileutil.Dir {
	result, err := fileutil.GetFile(rootpath)
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (app *FileApp) OpenfileByPath(path string) {
	fileutil.OpenfileByPath(path)
}

func (app *FileApp) DeletefileByPath(path string) {
	val, _ := config.GetConfigValue("DeleteToRecycleBin")
	if val != "false" {
		fileutil.MoveAll2Trash(path)
	} else {
		fileutil.DeletefileByPath(path)
	}
}

func (app *FileApp) RemoveAll2Trash(path string) {
	fileutil.MoveAll2Trash(path)
}

type RootpathInfo struct {
	Name     string
	Label    string
	Rootpath string
}

func init() {
	AddApp(NewFileApp())
}

func (app *FileApp) GetByRootpaths() []*RootpathInfo {
	var result []*RootpathInfo
	result = append(result, &RootpathInfo{
		Name:     "userfile",
		Label:    "用户根目录",
		Rootpath: fileutil.GetUserDir(),
	})
	result = append(result, &RootpathInfo{
		Name:     "desktop",
		Label:    "桌面",
		Rootpath: fileutil.GetDesktopDir(),
	})
	result = append(result, &RootpathInfo{
		Name:     "document",
		Label:    "文档",
		Rootpath: fileutil.GetMyDocumentDir(),
	})
	return result
}

func (app *FileApp) OpenDirectoryDialog() string {
	str, err := runtime.OpenDirectoryDialog(app.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		log.Println(err)
		return ""
	}
	return str
}
