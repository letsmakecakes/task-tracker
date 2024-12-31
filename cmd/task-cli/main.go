package main

import (
	"fmt"
	"log"
	"os"
	"tasktracker/internal/cli"
	"tasktracker/internal/storage"
)

func main() {
	// Get the file path from the environment variable.
	filePath := os.Getenv("TASKS_FILE_PATH")
	if filePath == "" {
		// Default to a specific path if the environment variable is not set.
		filePath = "./data/tasks.json"
	}

	// Initialize storage with the JSON file path
	store := storage.NewJSONStorage(filePath)

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
