package file

import (
	"fmt"
	"testing"
)

func TestGetAllDir(t *testing.T) {
	result, err := GetAllDir(GetUserDir())
	if err != nil {
		t.Error(err)
	}
	for _, v := range result {
		fmt.Println(v.RelativePath)
	}
}
