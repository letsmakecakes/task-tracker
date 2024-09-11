package task

import "errors"

type TaskStatus int

const (
	Todo TaskStatus = iota
	InProgress
	Done
)

var statusNames = []string{
	"todo",
	"in-progress",
	"done",
}

// String returns the string representation of a TaskStatus
func (s TaskStatus) String() string {
	if s < Todo || s > Done {
		return "unknown"
	}

	return statusNames[s]
}

// ParseStatus converts a string to a TaskStatus enum
func ParseStatus(status string) (TaskStatus, error) {
	for i, name := range statusNames {
		if name == status {
			return TaskStatus(i), nil
		}
	}

	return TaskStatus(-1), errors.New("invalid task status")
}
