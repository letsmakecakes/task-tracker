# Task Tracker CLI

A command-line task management application built in Go. This CLI tool helps you track and manage your tasks with features for adding, updating, marking progress, and listing tasks in different states.

## Features

- ✅ Add, update, and delete tasks
- 📊 Track task status (todo, in-progress, done)
- 📝 List all tasks or filter by status
- 💾 Persistent storage using JSON
- 🔒 Thread-safe file operations
- ⚡ No external dependencies

## Installation

### Prerequisites
- Go 1.19 or higher

### Building from Source

1. Clone the repository
```bash
git clone https://github.com/letsmakecakes/task-tracker.git
cd task-tracker
```

2. Build the application
```bash
go build -o task-cli cmd/task-cli/main.go
```

3. Set the data file path
```bash
export TASKS_FILE_PATH=/path/to/custom/tasks.json
```

4. (Optional) Add to PATH
```bash
mv task-cli /usr/local/bin/
```

## Usage

### Adding Tasks
```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Updating Tasks
```bash
task-cli update 1 "Buy groceries and cook dinner"
# Output: Task 1 updated successfully
```

### Deleting Tasks
```bash
task-cli delete 1
# Output: Task 1 deleted successfully
```

### Managing Task Status
```bash
task-cli mark-in-progress 1
# Output: Task 1 marked as in-progress

task-cli mark-done 1
# Output: Task 1 marked as done
```

### Listing Tasks
```bash
# List all tasks
task-cli list

# List by status
task-cli list todo
task-cli list in-progress
task-cli list done
```

## Project Structure

```
task-tracker/
├── cmd/
│   └── task-cli/
│       └── main.go              # Entry point
├── internal/
│   ├── cli/
│   │   ├── commands.go          # Command handlers
│   │   └── parser.go            # CLI argument parser
│   ├── models/
│   │   └── task.go             # Task struct and methods
│   ├── storage/
│   │   ├── json_storage.go     # JSON file storage
│   │   └── storage.go          # Storage interface
│   └── utils/
│       ├── errors.go           # Custom errors
│       └── time.go             # Time utilities
├── go.mod
└── README.md
```

## Task Properties

Each task contains the following information:
- `id`: Unique identifier
- `description`: Task description
- `status`: Current status (todo/in-progress/done)
- `createdAt`: Creation timestamp
- `updatedAt`: Last update timestamp

## Design Decisions

1. **File Storage**: Uses JSON for persistence due to its simplicity and human-readability
2. **Thread Safety**: Implements mutex locks for concurrent file operations
3. **Interface-Based Design**: Storage interface allows for easy testing and alternate implementations
4. **No External Dependencies**: Uses only Go standard library for maintainability

## Error Handling

The application handles various error cases:
- Invalid commands or arguments
- Invalid task IDs
- File operation errors
- Concurrent access issues

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Future Improvements

- [ ] Add task priorities
- [ ] Add due dates
- [ ] Add task categories/tags
- [ ] Add search functionality
- [ ] Add task notes/comments
- [ ] Add data export/import
- [ ] Add task completion statistics

## Support

If you have any questions or suggestions, please open an issue in the repository.