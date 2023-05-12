package tool

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func TrimSuffix(dirPath string, suffix string) {

	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("Failed to read directory: %v\n", err)
		return
	}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			oldFilePath := filepath.Join(dirPath, fileInfo.Name())
			newFileName := strings.TrimSuffix(fileInfo.Name(), suffix)
			newFilePath := filepath.Join(dirPath, newFileName)

			err = os.Rename(oldFilePath, newFilePath)
			if err != nil {
				fmt.Printf("Failed to rename file: %v\n", err)
				return
			}

			fmt.Printf("Renamed file from %s to %s\n", oldFilePath, newFilePath)
		}
	}
}
