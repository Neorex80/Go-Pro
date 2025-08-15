package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get the directory to organize
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	// Check if directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", dir)
		os.Exit(1)
	}

	// Organize files
	err := organizeFiles(dir)
	if err != nil {
		fmt.Printf("Error organizing files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File organization complete!")
}

func organizeFiles(dir string) error {
	// Read directory contents
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Process each file
	for _, file := range files {
		// Skip directories
		if file.IsDir() {
			continue
		}

		// Get file extension
		ext := strings.ToLower(filepath.Ext(file.Name()))
		
		// Skip files without extensions
		if ext == "" {
			continue
		}

		// Remove the dot from extension
		ext = ext[1:]

		// Create directory for extension if it doesn't exist
		extDir := filepath.Join(dir, ext)
		if _, err := os.Stat(extDir); os.IsNotExist(err) {
			err := os.Mkdir(extDir, 0755)
			if err != nil {
				return err
			}
		}

		// Move file to extension directory
		oldPath := filepath.Join(dir, file.Name())
		newPath := filepath.Join(extDir, file.Name())
		
		fmt.Printf("Moving %s to %s\n", file.Name(), extDir)
		err := os.Rename(oldPath, newPath)
		if err != nil {
			return err
		}
	}

	return nil
}
