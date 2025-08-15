package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define command-line flags
	prefix := flag.String("prefix", "", "Prefix to add to filenames")
	suffix := flag.String("suffix", "", "Suffix to add to filenames")
	replace := flag.String("replace", "", "String to replace in filenames")
	with := flag.String("with", "", "String to replace with in filenames")
	directory := flag.String("dir", ".", "Directory to process")

	flag.Parse()

	// Check if directory exists
	if _, err := os.Stat(*directory); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", *directory)
		os.Exit(1)
	}

	// Process files in directory
	err := filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Get the directory and filename
		dir := filepath.Dir(path)
		filename := filepath.Base(path)

		// Skip the current executable if it's in the directory
		if strings.Contains(filename, os.Args[0]) {
			return nil
		}

		// Apply transformations
		newFilename := filename

		if *prefix != "" {
			newFilename = *prefix + newFilename
		}

		if *suffix != "" {
			ext := filepath.Ext(newFilename)
			name := strings.TrimSuffix(newFilename, ext)
			newFilename = name + *suffix + ext
		}

		if *replace != "" && *with != "" {
			newFilename = strings.ReplaceAll(newFilename, *replace, *with)
		}

		// Rename file if name changed
		if newFilename != filename {
			oldPath := filepath.Join(dir, filename)
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("Renaming: %s -> %s\n", filename, newFilename)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error processing files: %v\n", err)
		os.Exit(1)
	}
}
