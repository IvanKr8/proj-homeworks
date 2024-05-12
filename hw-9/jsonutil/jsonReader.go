package jsonutil

import (
	"io/ioutil"
	"path/filepath"
)

func ReadJSONFile(filename string) ([]byte, error) {
	dir := "C:/Users/user/Documents/Sites/projector/go-course/homeworks/hw-9/jsonutil"
	fullPath := filepath.Join(dir, filename)

	// Теперь читайте файл из fullPath
	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
