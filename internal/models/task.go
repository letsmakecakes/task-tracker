package models

import "time"

// TaskStatus represents the status of a task.
type TaskStatus string

const (
	// StatusTodo represents a task that is yet to be started.
	StatusTodo TaskStatus = "todo"

	// StatusInProgress represents a task that is currently being worked on.
	StatusInProgress TaskStatus = "in-progress"

	// StatusDone represents a task that has been completed.
	StatusDone TaskStatus = "done"
)

// Task represents a task with a description, status, and timestamps.
type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
