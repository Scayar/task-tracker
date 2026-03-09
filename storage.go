package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadTasks() []Task {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	if json.Unmarshal(data, &tasks) != nil {
		return []Task{}
	}
	return tasks
}

func saveTasks(tasks []Task) bool {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println(err)
		return false
	}
	if os.WriteFile("tasks.json", data, 0644) != nil {
		fmt.Println("can't write tasks.json")
		return false
	}
	return true
}
