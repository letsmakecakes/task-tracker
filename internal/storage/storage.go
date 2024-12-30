package storage

import (
	"tasktracker/internal/models"
)

// Storage defines an interface for task storage operations.
type Storage interface {
	// AddTask adds a new task with the given description and returns the created task.
	AddTask(description string) (*models.Task, error)

	// UpdateTask updates the description of the task with the specified ID.
	UpdateTask(id int, description string) error

	// DeleteTask deletes the task with the specified ID.
	DeleteTask(id int) error

	// MarkTaskStatus updates the status of the task with the specified ID.
	MarkTaskStatus(id int, status models.TaskStatus) error

	// ListAllTasks returns all tasks.
	ListAllTasks() ([]models.Task, error)

	// ListTaskByStatus returns tasks that match the specified status.
	ListTaskByStatus(status models.TaskStatus) ([]models.Task, error)
}
