package task

import "time"

type Task struct {
	ID          string    `json:"id" mapstructure:"id"`
	Description string    `json:"description" mapstructure:"description"`
	Status      string    `json:"status" mapstructure:"status"`
	CreatedAt   time.Time `json:"created_at" mapstructure:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" mapstructure:"updated_at"`
}
