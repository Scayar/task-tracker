# Task Tracker CLI

https://github.com/Scayar/task-tracker

**Project URL:** https://github.com/Scayar/task-tracker  
**Repository:** https://github.com/Scayar/task-tracker

CLI tool to track and manage tasks. Built with Go (standard library only).

## Run

```
go run .
go build -o task-cli.exe .
```

## Commands

```
task-cli add "description"
task-cli update <id> "description"
task-cli delete <id>
task-cli mark-in-progress <id>
task-cli mark-done <id>
task-cli list
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Project Structure

- `main.go` - CLI parsing, command dispatch
- `task.go` - Task operations (add, update, delete, mark, list)
- `storage.go` - JSON file read/write
- `README.txt` - Full implementation explanation (exam requirement)
