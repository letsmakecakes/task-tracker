package cli

import (
	"fmt"
	"strconv"
	"tasktracker/internal/models"
	"tasktracker/internal/storage"
	"tasktracker/internal/utils"
)

type Commander struct {
	store storage.Storage
}

// NewCommander initializes a new Commander with the provided storage.
func NewCommander(store storage.Storage) *Commander {
	return &Commander{store: store}
}

// Add adds a new task with the provided description.
func (c *Commander) Add(args []string) error {
	if len(args) < 1 {
		return utils.ErrInvalidArgs
	}

	task, err := c.store.AddTask(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	return nil
}

// Update updates the description of an existing task.
func (c *Commander) Update(args []string) error {
	if len(args) < 2 {
		return utils.ErrInvalidArgs
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return utils.ErrInvalidID
	}

	if err := c.store.UpdateTask(id, args[1]); err != nil {
		return err
	}

	fmt.Printf("Task %d updates successfully\n", id)
	return nil
}

// Delete removes a task by its ID.
func (c *Commander) Delete(args []string) error {
	if len(args) < 1 {
		return utils.ErrInvalidArgs
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return utils.ErrInvalidID
	}

	if err := c.store.DeleteTask(id); err != nil {
		return err
	}

	fmt.Printf("Task %d deleted successfully\n", id)
	return nil
}

// MarkStatus updates the status of an existing task.
func (c *Commander) MarkStatus(args []string, status models.TaskStatus) error {
	if len(args) < 1 {
		return utils.ErrInvalidArgs
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return utils.ErrInvalidID
	}

	if err := c.store.MarkTaskStatus(id, status); err != nil {
		return err
	}

	fmt.Printf("Task %d marked as %s\n", id, status)
	return nil
}

// The List displays all tasks or tasks filtered by status.
func (c *Commander) List(args []string) error {
	var tasks []models.Task
	var err error

	if len(args) == 0 {
		tasks, err = c.store.ListAllTasks()
	} else {
		status := models.TaskStatus(args[0])
		if !isValidStatus(status) {
			return utils.ErrInvalidStatus
		}
		tasks, err = c.store.ListTaskByStatus(status)
	}

	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Printf("No tasks found")
		return nil
	}

	for _, task := range tasks {
		fmt.Printf("[%d] %s (Status: %s) - Created: %s\n", task.ID, task.Description, task.Status, utils.FormatTime(task.CreatedAt))
	}

	return nil
}

// isValidStatus checks if the provided status is valid.
func isValidStatus(status models.TaskStatus) bool {
	return status == models.StatusTodo || status == models.StatusInProgress || status == models.StatusDone
}
