package cli

import (
	"tasktracker/internal/models"
	"tasktracker/internal/storage"
	"tasktracker/internal/utils"
)

// Parser is responsible for parsing and executing commands.
type Parser struct {
	commander *Commander
}

// NewParser initializes a new Parser with the provided storage.
func NewParser(store storage.Storage) *Parser {
	return &Parser{
		commander: NewCommander(store),
	}
}

// Parse parses the input arguments and executes the corresponding command.
func (p *Parser) Parse(args []string) error {
	if len(args) < 1 {
		return utils.ErrInvalidCommand
	}

	command := args[0]
	cmdArgs := args[1:]

	switch command {
	case "add":
		return p.commander.Add(cmdArgs)
	case "update":
		return p.commander.Update(cmdArgs)
	case "delete":
		return p.commander.Delete(cmdArgs)
	case "mark-in-progress":
		return p.commander.MarkStatus(cmdArgs, models.StatusInProgress)
	case "mark-done":
		return p.commander.MarkStatus(cmdArgs, models.StatusDone)
	case "list":
		return p.commander.List(cmdArgs)
	default:
		return utils.ErrInvalidCommand
	}
}
