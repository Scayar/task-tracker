package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("add needs a description")
			os.Exit(1)
		}
		addTask(os.Args[2])
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("update needs id and description")
			os.Exit(1)
		}
		id := atoi(os.Args[2])
		if !updateTask(id, os.Args[3]) {
			os.Exit(1)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("delete needs id")
			os.Exit(1)
		}
		id := atoi(os.Args[2])
		if !deleteTask(id) {
			os.Exit(1)
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("mark-in-progress needs id")
			os.Exit(1)
		}
		id := atoi(os.Args[2])
		if !markTask(id, "in-progress") {
			os.Exit(1)
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("mark-done needs id")
			os.Exit(1)
		}
		id := atoi(os.Args[2])
		if !markTask(id, "done") {
			os.Exit(1)
		}
	case "list":
		filter := ""
		if len(os.Args) >= 3 {
			filter = os.Args[2]
			if filter != "done" && filter != "todo" && filter != "in-progress" {
				fmt.Printf("bad filter: %s\n", filter)
				os.Exit(1)
			}
		}
		listTasks(filter)
	default:
		fmt.Printf("unknown: %s\n", cmd)
		usage()
		os.Exit(1)
	}
}

func atoi(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("invalid id: %s\n", s)
		os.Exit(1)
	}
	return id
}

func usage() {
	fmt.Println("add <desc>")
	fmt.Println("update <id> <desc>")
	fmt.Println("delete <id>")
	fmt.Println("mark-in-progress <id>")
	fmt.Println("mark-done <id>")
	fmt.Println("list [done|todo|in-progress]")
}
