package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const dataFile = "tasks.json"

func loadTasks() []Task {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Warning: could not parse tasks.json, starting fresh")
		return []Task{}
	}
	return tasks
}

func saveTasks(tasks []Task) bool {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error: failed to encode tasks -", err)
		return false
	}
	if err := os.WriteFile(dataFile, data, 0644); err != nil {
		fmt.Println("Error: could not write tasks.json -", err)
		return false
	}
	return true
}
