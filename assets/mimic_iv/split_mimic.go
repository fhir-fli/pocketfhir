package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// SplitNDJSON splits a large NDJSON file into smaller chunks
func SplitNDJSON(inputFile string, outputDir string, linesPerFile int) error {
	// Open the input NDJSON file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Extract the base name of the input file (without extension)
	baseFileName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))

	// Initialize reader
	scanner := bufio.NewScanner(file)

	// Variables to track the current chunk
	var currentFile *os.File
	var currentLineCount int
	var fileIndex int

	for scanner.Scan() {
		// If starting a new chunk, create a new file
		if currentLineCount == 0 {
			outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_%d.ndjson", baseFileName, fileIndex))
			currentFile, err = os.Create(outputFile)
			if err != nil {
				return fmt.Errorf("failed to create output file: %v", err)
			}
			defer currentFile.Close()
			log.Printf("Created new chunk: %s", outputFile)
		}

		// Write the current line to the file
		_, err := currentFile.WriteString(scanner.Text() + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to output file: %v", err)
		}

		// Increment the line count
		currentLineCount++

		// If we reach the limit, close the current file and start a new one
		if currentLineCount >= linesPerFile {
			currentFile.Close()
			currentLineCount = 0
			fileIndex++
		}
	}

	if scanner.Err() != nil {
		return fmt.Errorf("error reading file: %v", scanner.Err())
	}

	// Close the last file if it's still open
	if currentFile != nil {
		currentFile.Close()
	}

	return nil
}

func main() {
	// Input file, output directory, and lines per file
	inputFile := "ObservationChartevents.ndjson"
	outputDir := "./"
	linesPerFile := 20000 // Adjust based on your needs

	err := SplitNDJSON(inputFile, outputDir, linesPerFile)
	if err != nil {
		log.Fatalf("Error splitting NDJSON file: %v", err)
	}

	log.Println("NDJSON file split successfully!")
}
