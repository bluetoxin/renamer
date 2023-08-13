package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var rootDir string
	var pattern string
	flag.StringVar(&rootDir, "dir", ".", "Root directory to start renaming files")
	flag.StringVar(&pattern, "pattern", "default", "Renaming pattern: default | custom")
	flag.Parse()

	switch pattern {
	case "default":
		renameFunc = RenameDefault
	case "custom":
		renameFunc = RenameCustom
	default:
		fmt.Println("Invalid renaming pattern")
		return
	}

	err := filepath.Walk(rootDir, visit)
	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}

type Renamer func(string) string

var renameFunc Renamer

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.IsDir() {
		newName := renameFunc(info.Name())
		newPath := filepath.Join(filepath.Dir(path), newName)
		err = os.Rename(path, newPath)
		if err != nil {
			fmt.Println("Error renaming file:", err)
		}
	}

	return nil
}
