package fileUtil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateOutputFile(inputFileName string, format string, outDir string) (*os.File, string, error) {

	err := os.MkdirAll(outDir, os.ModePerm)

	if err != nil {

		return nil, "", fmt.Errorf("cannot create output directory: %w", err)

	}

	baseName := filepath.Base(inputFileName)

	nameWithoutExt := strings.TrimSuffix(baseName, filepath.Ext(baseName))

	ext := strings.TrimPrefix(format, ".")

	outputFileName := nameWithoutExt + "." + ext

	outputPath := filepath.Join(outDir, outputFileName)

	outputFile, err := os.Create(outputPath)

	if err != nil {

		return nil, "", fmt.Errorf("cannot create output file: %w", err)

	}

	return outputFile, outputPath, nil

}
