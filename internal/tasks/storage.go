package tasks

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

const taskFilePath = "../../json/tasks.json"

// LoadTasks loads tasks from the JSON file
func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(taskFilePath); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(taskFilePath)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// SaveTasks saves tasks to the JSON file
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(taskFilePath, data, 0644)
}

// AddTask adds a new task
func AddTask(description string) (Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return Task{}, err
	}

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	if err := SaveTasks(tasks); err != nil {
		return Task{}, err
	}

	return newTask, nil
}

// UpdateTask updates a task's description
func UpdateTask(id int, description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

// DeleteTask removes a task by ID
func DeleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

// MarkTask updates the task's status
func MarkTask(id int, status TaskStatus) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}
