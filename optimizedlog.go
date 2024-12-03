package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

// LogEntry for JSON output structure
type LogEntry struct {
	Message string `json:"message"`
}

func main() {
	// Open the zip file
	zipFileName := "Day-4_log_file_1GB.zip"
	zipFile, err := os.Open(zipFileName)
	if err != nil {
		fmt.Println("Error opening zip file:", err)
		return
	}
	defer zipFile.Close()

	// Read the zip file contents
	zipReader, err := zip.NewReader(zipFile, getZipFileSize(zipFile))
	if err != nil {
		fmt.Println("Error reading zip file:", err)
		return
	}
	// Process the first file inside the zip (assuming only one file)
	if len(zipReader.File) == 0 {
		fmt.Println("No files in the zip archive")
		return
	}

	zipFileContent := zipReader.File[0]
	file, err := zipFileContent.Open()
	if err != nil {
		fmt.Println("Error opening file inside zip:", err)
		return
	}
	defer file.Close()

	// Process the log file with concurrency
	processLogFileConcurrently(file)
}

// Processes the log file concurrently
func processLogFileConcurrently(file io.ReadCloser) {
	scanner := bufio.NewScanner(file)

	// Channels for distributing lines and collecting results
	linesChan := make(chan string, 100) // Buffered to avoid blocking
	resultsChan := make(chan string, 100)
	var wg sync.WaitGroup

	// Start worker goroutines
	numWorkers := 4
	fmt.Println(runtime.NumCPU())
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(linesChan, resultsChan, &wg)
	}

	// Start a separate goroutine to close resultsChan once workers are done
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Read lines from the file and send them to the lines channel
	go func() {
		for scanner.Scan() {
			linesChan <- scanner.Text()
		}
		close(linesChan) // Signal that no more lines will be sent
	}()

	// Collect and print results
	for result := range resultsChan {
		fmt.Println(result)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// Worker function to process lines
func worker(linesChan <-chan string, resultsChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	seen := make(map[string]struct{}) // Local duplicate tracker

	for line := range linesChan {
		// Skip duplicate lines
		if _, exists := seen[line]; exists {
			continue
		}
		seen[line] = struct{}{}

		// Convert line to JSON
		entry := LogEntry{Message: line}
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			continue
		}

		// Send the result to the results channel
		resultsChan <- string(entryJSON)
	}
}

// Gets the size of the zip file
func getZipFileSize(file *os.File) int64 {
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file size:", err)
		return 0
	}
	return stat.Size()
}
