# Task Tracker CLI

Task Tracker CLI is a command-line interface (CLI) tool designed to help users manage their tasks efficiently. This tool allows you to add, update, delete, and list tasks, and also track the status of tasks (e.g., To Do, In Progress, Done). The tasks are stored in a JSON file on your local filesystem.

## Features

- Add new tasks
- Update or delete tasks
- Mark tasks as `in-progress` or `done`
- List all tasks or filter by status (`done`, `in-progress`, `todo`)
- Task persistence in a JSON file

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or above)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/task-tracker-cli.git
   cd task-tracker-cli
   ```

2. Build the project:

   ```bash
   go build -o tasktracker ./cmd/tasktracker
   ```

3. Run the executable:

   ```bash
   ./tasktracker
   ```

### Usage

The Task Tracker CLI provides several commands to manage tasks. Below are examples of how to use each command.

#### Add a New Task

```bash
./tasktracker add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

#### Update an Existing Task

```bash
./tasktracker update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully
```

#### Delete a Task

```bash
./tasktracker delete 1
# Output: Task deleted successfully
```

#### Mark a Task as In Progress

```bash
./tasktracker mark-in-progress 1
# Output: Task marked as in progress
```

#### Mark a Task as Done

```bash
./tasktracker mark-done 1
# Output: Task marked as done
```

#### List All Tasks

```bash
./tasktracker list
# Output: List of all tasks
```

#### List Tasks by Status

```bash
./tasktracker list todo
# Output: List of tasks with status "todo"
```

```bash
./tasktracker list done
# Output: List of tasks with status "done"
```

```bash
./tasktracker list in-progress
# Output: List of tasks with status "in-progress"
```

### Project Structure

The project follows a typical Go project structure with the following key directories and files:

```
task-tracker-cli/
│
├── cmd/
│   └── tasktracker/
│       └── main.go          # Entry point for the CLI application
│
├── internal/
│   └── tasks/
│       ├── task.go          # Task struct and logic for managing tasks
│       ├── storage.go       # Functions for interacting with the JSON file
│
├── pkg/
│   └── cli/
│       ├── cli.go           # CLI commands and argument parsing
│
├── json/
│   └── tasks.json           # File to store tasks (created automatically)
│
├── go.mod                   # Go module file
└── README.md                # Project documentation
```

### Task Properties

Each task contains the following properties:

- `id`: Unique identifier for the task
- `description`: Description of the task
- `status`: The status of the task (`todo`, `in-progress`, `done`)
- `createdAt`: The creation date and time of the task
- `updatedAt`: The last update date and time of the task

### Error Handling

The CLI gracefully handles errors such as:
- Invalid task IDs
- Missing commands or arguments
- Invalid status transitions (e.g., marking a non-existent task as done)
