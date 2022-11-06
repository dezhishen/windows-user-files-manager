package fileutil

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

type Dir struct {
	Name         string
	AbsolutePath string
	RelativePath string
	FileInfo     os.FileInfo
	Children     []*Dir
	IsFile       bool
}

func GetUserDir() string {
	return os.Getenv("USERPROFILE")
}

func getReltivePath(path string, rootPath string) string {
	str := strings.TrimPrefix(path, rootPath)
	str = strings.TrimPrefix(str, "\\")
	return str
}

func GetFileWithMaxDeep(maxDeep uint8, rootPath string) ([]*Dir, error) {
	fileInfoList, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}
	var result []*Dir
	var deep uint8 = 0
	for _, fileInfo := range fileInfoList {
		var dir = &Dir{
			Name:         fileInfo.Name(),
			AbsolutePath: filepath.Join(rootPath, fileInfo.Name()),
			RelativePath: getReltivePath(filepath.Join(rootPath, fileInfo.Name()), GetUserDir()),
			FileInfo:     fileInfo,
		}
		if fileInfo.IsDir() {
			children, err := getFileWithMaxDeep(deep+1, maxDeep, dir.AbsolutePath)
			if err != nil {
				log.Printf("getFileWithMaxDeep error: %v", err)
			}
			if children != nil {
				dir.Children = append(dir.Children, children...)
			}
		}
		result = append(result, dir)
	}
	return result, err
}

func getFileWithMaxDeep(deep, maxDeep uint8, rootPath string) ([]*Dir, error) {
	if deep > maxDeep {
		return nil, nil
	}
	fileInfoList, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}
	var result []*Dir
	for _, fileInfo := range fileInfoList {
		var dir = &Dir{
			Name:         fileInfo.Name(),
			AbsolutePath: filepath.Join(rootPath, fileInfo.Name()),
			RelativePath: getReltivePath(filepath.Join(rootPath, fileInfo.Name()), GetUserDir()),
			FileInfo:     fileInfo,
		}
		if fileInfo.IsDir() {
			children, err := getFileWithMaxDeep(deep+1, maxDeep, dir.AbsolutePath)
			if err != nil {
				continue
			}
			if children != nil {
				dir.Children = append(dir.Children, children...)
			}
		}
		result = append(result, dir)
	}
	return result, err
}

func GetFile(rootPath string) ([]*Dir, error) {
	if rootPath == "" {
		rootPath = GetUserDir()
	}
	fileInfoList, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}
	var result []*Dir
	for _, fileInfo := range fileInfoList {
		absolutePath := filepath.Join(rootPath, fileInfo.Name())
		var dir = &Dir{
			Name:         fileInfo.Name(),
			AbsolutePath: absolutePath,
			RelativePath: getReltivePath(filepath.Join(rootPath, fileInfo.Name()), rootPath),
			FileInfo:     fileInfo,
			IsFile:       !fileInfo.IsDir(),
		}
		result = append(result, dir)
	}
	return result, err
}

func OpenfileByPath(path string) {
	if runtime.GOOS == "darwin" {
		cmd := `open "` + path + `"`
		exec.Command("/bin/bash", "-c", cmd).Start()
	} else {
		path = strings.ReplaceAll(path, "/", "\\")
		exec.Command("explorer", path).Start()
	}
}

func GetMyDocumentDir() string {
	return getFolderPath("MyDocuments")
}

func GetDesktopDir() string {
	return getFolderPath("Desktop")
}

func getFolderPath(str string) string {
	back := &backend.Local{}
	// 开启一个本地 powershell 进程
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	// ... 和它交互
	cmd := fmt.Sprintf("[Environment]::GetFolderPath(\"%s\")", str)
	stdout, _, err := shell.Execute(cmd)
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(strings.TrimSpace(stdout), "\\", "\\\\")
}

func DeletefileByPath(path string) error {
	return os.RemoveAll(path)
}
