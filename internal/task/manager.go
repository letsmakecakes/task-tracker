package task

import (
	"encoding/json"
	"errors"
	"os"
)

// TaskManager manages a list of tasks and interacts with the JSON file
type TaskManager struct {
	Tasks      []*Task
	filePath   string
	lastTaskID int
}

// NewTaskManager initializes a new TaskManager and loads tasks from a JSON file
func NewTaskManager(filePath string) (*TaskManager, error) {
	manager := &TaskManager{
		filePath: filePath,
	}

	if err := manager.loadTasks(); err != nil {
		return nil, err
	}

	return manager, nil
}

// loadTasks loads tasks from the JSON file, or creates a new file if it doesn't exist
func (m *TaskManager) loadTasks() error {
	if _, err := os.Stat(m.filePath); errors.Is(err, os.ErrNotExist) {
		return m.saveTasks() // Create file if it doesn't exist
	}

	data, err := os.ReadFile(m.filePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &m.Tasks); err != nil {
		return err
	}

	// Find the last task ID
	for _, task := range m.Tasks {
		if task.ID > m.lastTaskID {
			m.lastTaskID = task.ID
		}
	}

	return nil
}

// saveTasks saves the current list of tasks to the JSON file
func (m *TaskManager) saveTasks() error {
	data, err := json.MarshalIndent(m.Tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(m.filePath, data, 0644)
}

// AddTask adds a new task to the list and saves it to the JSON file
func (m *TaskManager) AddTask(description string) (*Task, error) {
	m.lastTaskID++
	newTask := NewTask(m.lastTaskID, description)
	m.Tasks = append(m.Tasks, newTask)

	if err := m.saveTasks(); err != nil {
		return nil, err
	}

	return newTask, nil
}

// UpdateTask updates an existing task and saves the changes to the JSON file
func (m *TaskManager) UpdateTask(id int, description string, status TaskStatus) error {
	task, err := m.findTaskByID(id)
	if err != nil {
		return err
	}

	task.UpdateTask(description, status)
	return m.saveTasks()
}

// DeleteTask deletes a task by its ID and saves the changes to the JSON file
func (m *TaskManager) DeleteTask(id int) error {
	for i, task := range m.Tasks {
		if task.ID == id {
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			return m.saveTasks()
		}
	}
	return errors.New("task not found")
}

// ListTasks lists all tasks
func (m *TaskManager) ListTasks() []*Task {
	return m.Tasks
}

// ListTasksByStatus lists tasks filtered by their status
func (m *TaskManager) ListTasksByStatus(status TaskStatus) []*Task {
	var filteredTasks []*Task
	for _, task := range m.Tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

// MarkTaskInProgress marks a task as in progress
func (m *TaskManager) MarkTaskInProgress(id int) error {
	task, err := m.findTaskByID(id)
	if err != nil {
		return err
	}

	task.MarkInProgress()
	return m.saveTasks()
}

// MarkTaskDone marks a task as done
func (m *TaskManager) MarkTaskDone(id int) error {
	task, err := m.findTaskByID(id)
	if err != nil {
		return err
	}

	task.MarkDone()
	return m.saveTasks()
}

// findTaskByID finds a task by its ID
func (m *TaskManager) findTaskByID(id int) (*Task, error) {
	for _, task := range m.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, errors.New("task not found")
}
