package cli

import (
	"strconv"
	"tasktracker/internal/tasks"

	log "github.com/sirupsen/logrus"
)

func HandleCommand(args []string) {
	switch args[0] {
	case "add":
		if len(args) < 2 {
			log.Error("please provide a task description")
			return
		}
		task, err := tasks.AddTask(args[1])
		if err != nil {
			log.Errorf("error adding task: %v", err)
		} else {
			log.Printf("task added successfully (ID: %d)\n", task.ID)
		}

	case "list":
		if len(args) == 2 {
			filterStatus := args[1]
			ListTaskByStatus(filterStatus)
		} else {
			ListAllTasks()
		}

	case "update":
		if len(args) < 3 {
			log.Error("usage: update <id> <new description>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			log.Errorf("invalid ID: %v", args[1])
			return
		}
		if err := tasks.UpdateTask(id, args[2]); err != nil {
			log.Errorf("error updating task: %v", err)
		} else {
			log.Println("task updated successfully")
		}

	case "delete":
		if len(args) < 2 {
			log.Error("usage: delete <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			log.Errorf("invalid ID: %v", args[1])
			return
		}
		if err := tasks.DeleteTask(id); err != nil {
			log.Errorf("error deleting task: %v", err)
		} else {
			log.Println("task deleted successfully")
		}
	case "mark-in-progress":
		if len(args) < 2 {
			log.Error("usage: mark-in-progress <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			log.Errorf("invalid ID: %v", args[1])
			return
		}
		if err := tasks.MarkTask(id, tasks.InProgress); err != nil {
			log.Errorf("error marking task in progress: %v", err)
		} else {
			log.Println("task marked as in-progress")
		}
	case "mark-done":
		if len(args) < 2 {
			log.Errorf("Usage: mark-done <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			log.Errorf("invalid ID: %v", args[1])
			return
		}
		if err := tasks.MarkTask(id, tasks.Done); err != nil {
			log.Errorf("error marking task done: %v", err)
		} else {
			log.Println("task marked as done")
		}
	default:
		log.Errorf("Unknown command: %v", args[0])
		log.Error("Available commands: add, list, update, delete, mark-in-progress, mark-done")
	}
}

// ListAllTasks prints all tasks
func ListAllTasks() {
	taskList, err := tasks.LoadTasks()
	if err != nil {
		log.Errorf("error loading tasks: %v", err)
		return
	}

	for _, task := range taskList {
		log.Printf("%d: %s [%s] created at: %v last updated at: %v", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

// ListTasksByStatus filters and prints tasks by status
func ListTaskByStatus(status string) {
	taskList, err := tasks.LoadTasks()
	if err != nil {
		log.Errorf("error loading tasks: %v", err)
		return
	}

	for _, task := range taskList {
		if string(task.Status) == status {
			log.Printf("%d: %s [%s] created at: %v last updated at: %v", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}
