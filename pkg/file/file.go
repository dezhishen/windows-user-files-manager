package file

import (
	"os"
	"path/filepath"
	"strings"
)

type Dir struct {
	Name         string
	AbsolutePath string
	RelativePath string
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
		})
		return nil
	})
	return result, err
}

func getReltivePath(path string, rootPath string) string {
	str := strings.TrimPrefix(path, rootPath)
	str = strings.Replace(str, "\\", "/", -1)
	str = strings.TrimPrefix(str, "/")
	return str
}
