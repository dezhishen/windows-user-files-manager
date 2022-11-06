package application

import (
	"context"
	"log"

	"github.com/dezhishen/windows-user-files-manager/pkg/fileutil"
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
func (a *FileApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *FileApp) GetAllDir() []*fileutil.Dir {
	result, err := fileutil.GetFile(fileutil.GetUserDir())
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (a *FileApp) GetByRootpath(rootpath string) []*fileutil.Dir {
	result, err := fileutil.GetFile(rootpath)
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (a *FileApp) OpenfileByPath(path string) {
	fileutil.OpenfileByPath(path)
}
