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

func addTask(desc string) {
	tasks := loadTasks()
	id := 1
	for _, t := range tasks {
		if t.ID >= id {
			id = t.ID + 1
		}
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	tasks = append(tasks, Task{id, desc, "todo", now, now})
	if !saveTasks(tasks) {
		os.Exit(1)
	}
	fmt.Printf("%d\n", id)
}

func updateTask(id int, desc string) bool {
	tasks := loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return saveTasks(tasks)
		}
	}
	fmt.Printf("task %d not found\n", id)
	return false
}

func deleteTask(id int) bool {
	tasks := loadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return saveTasks(tasks)
		}
	}
	fmt.Printf("task %d not found\n", id)
	return false
}

func markTask(id int, status string) bool {
	tasks := loadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return saveTasks(tasks)
		}
	}
	fmt.Printf("task %d not found\n", id)
	return false
}

func listTasks(filter string) {
	tasks := loadTasks()
	for _, t := range tasks {
		if filter != "" && t.Status != filter {
			continue
		}
		icon := "[ ]"
		if t.Status == "done" {
			icon = "[x]"
		} else if t.Status == "in-progress" {
			icon = "[~]"
		}
		fmt.Printf("%s %d %s\n", icon, t.ID, t.Description)
	}
}
