package io

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Check current file path is exist
func GetAbsolutelyFilePath(path string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// If input file path is not absolutely filepath, let's join it with current working directory
	if !filepath.IsAbs(path) {
		path = filepath.Join(wd, strings.Trim(path, "."))
	}

	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	//
	if info.IsDir() {
		return "", fmt.Errorf("%v is is not a file", path)
	}

	return path, nil
}
