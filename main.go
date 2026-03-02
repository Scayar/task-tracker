package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		requireArgs(3, "task-cli add \"Buy groceries\"")
		addTask(os.Args[2])

	case "update":
		requireArgs(4, "task-cli update 1 \"New description\"")
		id := parseID(os.Args[2])
		if !updateTask(id, os.Args[3]) {
			os.Exit(1)
		}

	case "delete":
		requireArgs(3, "task-cli delete 1")
		id := parseID(os.Args[2])
		if !deleteTask(id) {
			os.Exit(1)
		}

	case "mark-in-progress":
		requireArgs(3, "task-cli mark-in-progress 1")
		id := parseID(os.Args[2])
		if !markTask(id, "in-progress") {
			os.Exit(1)
		}

	case "mark-done":
		requireArgs(3, "task-cli mark-done 1")
		id := parseID(os.Args[2])
		if !markTask(id, "done") {
			os.Exit(1)
		}

	case "list":
		filter := ""
		if len(os.Args) >= 3 {
			filter = os.Args[2]
			if filter != "done" && filter != "todo" && filter != "in-progress" {
				fmt.Printf("Unknown filter: %s\n", filter)
				os.Exit(1)
			}
		}
		listTasks(filter)

	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func requireArgs(n int, example string) {
	if len(os.Args) < n {
		fmt.Printf("Not enough arguments. Example: %s\n", example)
		os.Exit(1)
	}
}

func parseID(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("\"%s\" is not a valid task ID\n", s)
		os.Exit(1)
	}
	return id
}

func printUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  add \"description\"            Add a new task")
	fmt.Println("  update <id> \"description\"    Update task description")
	fmt.Println("  delete <id>                  Delete a task")
	fmt.Println("  mark-in-progress <id>        Mark task as in-progress")
	fmt.Println("  mark-done <id>               Mark task as done")
	fmt.Println("  list                         List all tasks")
	fmt.Println("  list done                    List completed tasks")
	fmt.Println("  list todo                    List pending tasks")
	fmt.Println("  list in-progress             List in-progress tasks")
}
