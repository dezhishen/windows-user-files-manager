package application

import (
	"context"

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
func (a *FileApp) GetAllDir() []fileutil.Dir {
	var result []fileutil.Dir
	fileutil.ScanFile(
		fileutil.GetUserDir(),
		func(dir *fileutil.Dir) error {
			result = append(result, *dir)
			return nil
		},
	)
	return result
}
