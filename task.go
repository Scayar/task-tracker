package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func findTask(tasks []Task, id int) int {
	for i, t := range tasks {
		if t.ID == id {
			return i
		}
	}
	return -1
}

func addTask(desc string) {
	tasks := loadTasks()

	task := Task{
		ID:          nextID(tasks),
		Description: desc,
		Status:      "todo",
		CreatedAt:   timestamp(),
		UpdatedAt:   timestamp(),
	}

	tasks = append(tasks, task)
	if !saveTasks(tasks) {
		os.Exit(1)
	}
	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

func updateTask(id int, newDesc string) bool {
	tasks := loadTasks()
	idx := findTask(tasks, id)
	if idx == -1 {
		fmt.Printf("Error: task %d not found\n", id)
		return false
	}
	tasks[idx].Description = newDesc
	tasks[idx].UpdatedAt = timestamp()
	if !saveTasks(tasks) {
		os.Exit(1)
	}
	fmt.Println("Task updated")
	return true
}

func deleteTask(id int) bool {
	tasks := loadTasks()
	idx := findTask(tasks, id)
	if idx == -1 {
		fmt.Printf("Error: task %d not found\n", id)
		return false
	}
	tasks = append(tasks[:idx], tasks[idx+1:]...)
	if !saveTasks(tasks) {
		os.Exit(1)
	}
	fmt.Println("Task deleted")
	return true
}

func markTask(id int, status string) bool {
	tasks := loadTasks()
	idx := findTask(tasks, id)
	if idx == -1 {
		fmt.Printf("Error: task %d not found\n", id)
		return false
	}
	tasks[idx].Status = status
	tasks[idx].UpdatedAt = timestamp()
	if !saveTasks(tasks) {
		os.Exit(1)
	}
	fmt.Printf("Task %d marked as %s\n", id, status)
	return true
}

func listTasks(filter string) {
	tasks := loadTasks()

	if len(tasks) == 0 {
		fmt.Println("No tasks yet. Add one with: task-cli add \"your task\"")
		return
	}

	count := 0
	for _, t := range tasks {
		if filter != "" && t.Status != filter {
			continue
		}
		statusTag := statusIcon(t.Status)
		fmt.Printf("  %s [%d] %s\n", statusTag, t.ID, t.Description)
		count++
	}

	if count == 0 {
		fmt.Printf("No tasks with status \"%s\"\n", filter)
	} else {
		fmt.Printf("\nShowing %d task(s)\n", count)
	}
}

func statusIcon(status string) string {
	switch status {
	case "done":
		return "[x]"
	case "in-progress":
		return "[~]"
	default:
		return "[ ]"
	}
}
