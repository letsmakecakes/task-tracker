package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"tasktracker/internal/models"
	"time"
)

// JSONStorage is a struct that provides thread-safe access to a JSON file storing tasks.
type JSONStorage struct {
	filename string
	mutex    sync.RWMutex
}

// NewJSONStorage initializes a new JSONStorage with the given filename.
func NewJSONStorage(filename string) *JSONStorage {
	return &JSONStorage{filename: filename}
}

// loadTasks loads storage from the JSON file
func (s *JSONStorage) loadTasks() ([]models.Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Create the file if it doesn't exist
	if _, err := os.Stat(s.filename); os.IsNotExist(err) {
		return []models.Task{}, nil
	}

	data, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// saveTasks saves storage to the JSON file
func (s *JSONStorage) saveTasks(tasks []models.Task) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, data, 0644)
}

// AddTask adds a new task with the given description.
func (s *JSONStorage) AddTask(description string) (*models.Task, error) {
	tasks, err := s.loadTasks()
	if err != nil {
		return nil, err
	}

	newTask := models.Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      models.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	if err := s.saveTasks(tasks); err != nil {
		return nil, err
	}

	return &newTask, nil
}

// UpdateTask updates the description of a task by its ID.
func (s *JSONStorage) UpdateTask(id int, description string) error {
	tasks, err := s.loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return s.saveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

// DeleteTask removes a task by its ID.
func (s *JSONStorage) DeleteTask(id int) error {
	tasks, err := s.loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return s.saveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

// MarkTaskStatus updates the status of a task by its ID.
func (s *JSONStorage) MarkTaskStatus(id int, status models.TaskStatus) error {
	tasks, err := s.loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return s.saveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

// ListAllTasks returns all tasks.
func (s *JSONStorage) ListAllTasks() ([]models.Task, error) {
	return s.loadTasks()
}

// ListTaskByStatus returns tasks that match the specified status.
func (s *JSONStorage) ListTaskByStatus(status models.TaskStatus) ([]models.Task, error) {
	tasks, err := s.loadTasks()
	if err != nil {
		return nil, err
	}

	var filteredTasks []models.Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks, nil
}
