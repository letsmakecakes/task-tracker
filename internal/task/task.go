package task

import "time"

type Task struct {
	ID          int        `json:"id" mapstructure:"id"`
	Description string     `json:"description" mapstructure:"description"`
	Status      TaskStatus `json:"status" mapstructure:"status"`
	CreatedAt   time.Time  `json:"created_at" mapstructure:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" mapstructure:"updated_at"`
}

// NewTask creates a new task with the given descriptuon and default status
func NewTask(id int, description string) *Task {
	now := time.Now()
	return &Task{
		ID:          id,
		Description: description,
		Status:      Todo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateTask updates the description and status of the a task
func (t *Task) UpdateTask(description string, status TaskStatus) {
	t.Description = description
	t.Status = status
	t.UpdatedAt = time.Now()
}

// MarkInProgress marks the task as "in-progress"
func (t *Task) MarkInProgress() {
	t.Status = InProgress
	t.UpdatedAt = time.Now()
}

// MarkDone marks the task as "done"
func (t *Task) MarkDone() {
	t.Status = Done
	t.UpdatedAt = time.Now()
}
