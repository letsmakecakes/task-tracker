package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"tasktracker/internal/cli"
	"tasktracker/internal/storage"
)

func main() {
	// Get the directory of the executable
	execDir, err := os.Executable()
	if err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if err2 != nil {
			log.Fatalf("error printing message to console: %v", err2)
		}
		os.Exit(1)
	}

	// Build the path to the JSON file in the data directory
	dataFilePath := filepath.Join(filepath.Dir(execDir), "data", "tasks.json")

	// Initialize storage with the JSON file path
	store := storage.NewJSONStorage(dataFilePath)

	// Create a new command parser with the storage.
	parser := cli.NewParser(store)

	// Parse the command-line arguments and execute the appropriate command.
	if err := parser.Parse(os.Args[1:]); err != nil {
		// Print the error message to the standard error stream.
		_, err2 := fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if err2 != nil {
			log.Fatalf("error printing message to console: %v", err2)
		}
		// Exit with a non-zero status code to indicate an error.
		os.Exit(1)
	}
}
