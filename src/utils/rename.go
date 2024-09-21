package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func Rename(directoryName string) {
	directory, err := os.Open(directoryName)
	if err != nil {
		panic(err)
	}
	defer directory.Close()

	files, err := directory.Readdir(-1)
	if err != nil {
		panic(err)
	}

	var jpgFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".jpg") {
			jpgFiles = append(jpgFiles, file.Name())
		}
	}

	sort.Strings(jpgFiles)

	for i, oldName := range jpgFiles {
		newName := strconv.Itoa(i+1) + ".jpg"
		oldPath := filepath.Join(directoryName, oldName)
		newPath := filepath.Join(directoryName, newName)

		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Error renaming %s to %s: %v\n", oldName, newName, err)
		}
	}
}
