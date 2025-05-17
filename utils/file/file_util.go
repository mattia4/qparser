package fileUtil

import (
	"os"
	"path/filepath"
	"strings"
)

func CreateOutputFile(inputFileName string, ext string, dirName string) (*os.File, string, error) {

	fileName := strings.TrimSuffix(filepath.Base(inputFileName), filepath.Ext(inputFileName)) + ext

	outputPath := filepath.Join(dirName, fileName)

	f, err := createFile(outputPath)

	return f, outputPath, err

}

func createFile(name string) (*os.File, error) {

	_, e := os.Stat(name)

	if e == nil {
		return nil, os.ErrExist
	}

	return os.Create(name)
}
