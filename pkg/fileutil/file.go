package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

type Dir struct {
	Name         string
	AbsolutePath string
	RelativePath string
	FileInfo     os.FileInfo
}

func GetUserDir() string {
	return os.Getenv("USERPROFILE")
}

//
func GetAllDir(rootPath string) ([]*Dir, error) {
	result := make([]*Dir, 0)
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		result = append(result, &Dir{
			Name:         info.Name(),
			AbsolutePath: path,
			RelativePath: getReltivePath(path, rootPath),
			FileInfo:     info,
		})
		return nil
	})
	return result, err
}

func getReltivePath(path string, rootPath string) string {
	str := strings.TrimPrefix(path, rootPath)
	// str = strings.Replace(str, "\\", "/", -1)
	str = strings.TrimPrefix(str, "\\")
	return str
}

func ScanFile(rootPath string, callback func(*Dir) error) error {
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		dir := &Dir{
			Name:         info.Name(),
			AbsolutePath: path,
			RelativePath: getReltivePath(path, rootPath),
			FileInfo:     info,
		}
		theErr := callback(dir)
		if theErr != nil {
			return theErr
		}
		return nil
	})
	return err
}
