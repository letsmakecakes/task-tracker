package main

import (
	"os"
	"tasktracker/pkg/cli"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Get the command and arguments
	if len(os.Args) < 2 {
		log.Panic("expected a command (add, list update, delete, mark-done, mark-in-progress)")
	}

	// Parse and execute the command
	cli.HandleCommand(os.Args[1:])
}
