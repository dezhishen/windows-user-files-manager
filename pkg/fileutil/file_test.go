package fileutil

import (
	"fmt"
	"testing"
)

func TestGetAllDir(t *testing.T) {
	result, err := GetFileWithMaxDeep(2, GetUserDir())
	if err != nil {
		t.Error(err)
	}
	for _, v := range result {
		printDir(v, 0)
	}
}

func printDir(dir *Dir, deep uint8) {
	for _, v := range dir.Children {
		tab := ""
		for i := 0; i < int(deep); i++ {
			tab += "-"
		}
		fmt.Println(tab + v.Name)
		if v.Children != nil {
			for _, child := range v.Children {
				printDir(child, deep+1)
			}
		}
	}
}
