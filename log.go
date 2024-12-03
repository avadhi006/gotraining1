package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// LogEntry for JSON output structure
type LogEntry struct {
	Message string `json:"message"`
}

func main() {
	//extract zip
	zipFileName := "Day-4_log_file_1GB.zip"
	zipFile, err := os.Open(zipFileName)
	if err != nil {
		fmt.Println("Error opening zip file:", err)
		return
	}
	defer zipFile.Close()

	// Read the zip file's contents
	zipReader, err := zip.NewReader(zipFile, getZipFileSize(zipFile))
	if err != nil {
		fmt.Println("Error reading zip file:", err)
		return
	}

	// Iterate through the files in the zip 
	for _, zipFile := range zipReader.File {
		// Open each file inside the zip 
		file, err := zipFile.Open()
		if err != nil {
			fmt.Println("Error opening file inside zip:", err)
			continue
		}
		defer file.Close()

		// Processing the file 
		fmt.Printf("Processing file: %s\n", zipFile.Name)
		processLogFile(file)
	}
}

// read logfile and converting each line into a JSON object
func processLogFile(file io.ReadCloser) {
	scanner := bufio.NewScanner(file)

	// track of duplicates
	seen := make(map[string]struct{})
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		// Print debugging information for each line
		fmt.Printf("Processing line %d: %s\n", lineCount, line)

		
		if _, exists := seen[line]; exists {
			// Skip duplicate line
			fmt.Printf("Skipping duplicate line %d: %s\n", lineCount, line)
			continue
		}
		// Mark duplicate line as seen
		seen[line] = struct{}{}

		// Create a LogEntry struct with the message
		entry := LogEntry{
			Message: line,
		}

		// Convert the log entry into JSON
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			continue
		}

		// Print the JSON representation of the log entry
		fmt.Println(string(entryJSON))
	}

	// Check for any errors during scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func getZipFileSize(file *os.File) int64 {
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file size:", err)
		return 0
	}
	return stat.Size()
}
