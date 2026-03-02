Task Tracker CLI - Implementation Explanation
============================================

Project Page URL: https://roadmap.sh/projects/task-tracker
Repository: https://github.com/Scayar/task-tracker

This document explains exactly how the Task Tracker CLI application was programmed.
Language: Go. No external libraries - only Go standard library (os, fmt, strconv, encoding/json, time).


HOW TO RUN
----------
  go run .
  go build -o task-cli.exe .


COMMANDS (positional arguments as required)
-------------------------------------------
  task-cli add "description"
  task-cli update <id> "description"
  task-cli delete <id>
  task-cli mark-in-progress <id>
  task-cli mark-done <id>
  task-cli list
  task-cli list done
  task-cli list todo
  task-cli list in-progress


PROJECT STRUCTURE
-----------------
  main.go    - CLI entry point, parses os.Args and dispatches to task functions
  task.go    - Task struct, all task operations (add, update, delete, mark, list)
  storage.go - JSON file read/write using native os and encoding/json
  tasks.json - Created automatically in current directory when first task is added
  go.mod     - Go module (no external dependencies)


HOW IT WAS PROGRAMMED
---------------------

1. main.go - Command Line Interface
   - Reads user input from os.Args (positional arguments)
   - First argument = command (add, update, delete, mark-in-progress, mark-done, list)
   - Validates argument count with requireArgs() before each operation
   - Uses parseID() to convert string ID to int, exits with error if invalid
   - For "list", optional second argument = filter (done, todo, in-progress)
   - Unknown commands or filters print error and show usage
   - On task-not-found (update/delete/mark), exits with code 1

2. task.go - Task Logic
   - Task struct: id, description, status, createdAt, updatedAt (all required properties)
   - nextID(): finds max ID in list + 1 for new tasks
   - findTask(): returns index by id, -1 if not found
   - addTask(): creates task with status "todo", timestamps via time.Now()
   - updateTask(): updates description and updatedAt
   - deleteTask(): removes task by index with slice append
   - markTask(): sets status to "in-progress" or "done", updates updatedAt
   - listTasks(): filters by status when given, shows [x] done, [~] in-progress, [ ] todo

3. storage.go - JSON File (native filesystem)
   - dataFile constant = "tasks.json" in current directory
   - loadTasks(): os.ReadFile - if file missing, returns empty slice (file created on first save)
   - loadTasks(): json.Unmarshal - if parse fails, prints warning and returns empty slice
   - saveTasks(): json.MarshalIndent for readable JSON, os.WriteFile - creates file if not exists
   - saveTasks() returns false on error; callers exit with 1 to avoid false success message


ERROR HANDLING
--------------
  - Missing/invalid arguments: print example, exit 1
  - Invalid task ID (non-numeric): print error, exit 1
  - Task not found: print error, exit 1
  - Unknown filter: print error, exit 1
  - JSON write/encode failure: print error, exit 1
  - Corrupt tasks.json: warning, start with empty list


JSON FILE FORMAT
----------------
  [
    {
      "id": 1,
      "description": "Buy groceries",
      "status": "todo",
      "createdAt": "2026-03-02 14:00:00",
      "updatedAt": "2026-03-02 14:00:00"
    }
  ]


REQUIREMENTS MET
----------------
  - Add, Update, Delete tasks
  - Mark in-progress, mark done
  - List all / list done / list todo / list in-progress
  - Positional arguments
  - JSON file in current directory, created if not exists
  - Native filesystem (os package)
  - No external libraries
  - Error handling and edge cases
