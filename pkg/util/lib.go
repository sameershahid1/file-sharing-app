package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateFile(folderPath string, fileName string) (*os.File, error) {
	// Ensure the directory exists (creates if it doesn't)
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// Construct the full file path
	filePath := filepath.Join(folderPath, fileName)

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	log.Printf("File successfully created at: %s", filePath)
	return file, nil
}





