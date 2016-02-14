package main

import (
	"os"
	"path/filepath"
	"strings"
)

func initFolders() {
	runnerLog("InitFolders")
	path := tmpPath()
	runnerLog("mkdir %s", path)
	err := os.Mkdir(path, 0755)
	if err != nil {
		runnerLog(err.Error())
	}
}

func isTmpDir(path string) bool {
	absolutePath, _ := filepath.Abs(path)
	absoluteTmpPath, _ := filepath.Abs(tmpPath())

	return absolutePath == absoluteTmpPath
}

func isWatchedFile(path string) bool {
	absolutePath, _ := filepath.Abs(path)
	absoluteTmpPath, _ := filepath.Abs(tmpPath())

	if strings.HasPrefix(absolutePath, absoluteTmpPath) {
		return false
	}

	ext := filepath.Ext(path)

	for _, e := range strings.Split(settings["valid_ext"], ",") {
		if strings.TrimSpace(e) == ext {
			return true
		}
	}

	return false
}

func createBuildErrorsLog(message string) bool {
	file, err := os.Create(buildErrorsFilePath())
	if err != nil {
		return false
	}

	_, err = file.WriteString(message)
	if err != nil {
		return false
	}

	return true
}

func removeBuildErrorsLog() error {
	err := os.Remove(buildErrorsFilePath())

	return err
}
