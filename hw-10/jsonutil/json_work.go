package jsonutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func OpenJSONFile(filename string) ([]byte, error) {
	dir := "./jsonutil"
	fullPath := filepath.Join(dir, filename)

	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SaveJSONFile(filename string, data []byte) error {
	dir := "./jsonutil"
	fullPath := filepath.Join(dir, filename)

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
