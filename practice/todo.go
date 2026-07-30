package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const fileName string = "todo.json"

func main() {

	var err error

	if len(os.Args) == 3 {
		err = run(os.Args[1], os.Args[2])
	}
	if len(os.Args) == 2 {
		err = run(os.Args[1], "")
	}

	if err != nil {
		fmt.Println(err)
	}
}

func run(cmd string, data string) error {
	switch cmd {
	case "list":
		data, err := readTasks()
		if err != nil {
			return fmt.Errorf("Error saving the Tasks file")
		}
		printTasks(data)
	case "add":
		return saveTask(Task{Task: data, Done: false})
	case "done":
		id, _ := strconv.Atoi(data)
		return markDone(id)
	}

	return nil

}
func printTasks(data []Task) {
	for _, task := range data {
		fmt.Printf("Id: %d\nTask: %s\nDone: %t\n\n", task.ID, task.Task, task.Done)
	}
}

func markDone(id int) error {
	data, err := readTasks()
	emptyFile()
	if err != nil {
		return fmt.Errorf("Error saving the Tasks file")
	}
	for _, task := range data {
		if task.ID == id {
			task.Done = true
		}
		saveTask(task)
	}
	return nil
}

func emptyFile() error {
	strData, err := json.MarshalIndent([]string{}, "", "  ")
	if err != nil {
		return fmt.Errorf("Error saving the Tasks file")
	}

	return os.WriteFile(fileName, strData, 0644)
}

func saveTask(task Task) error {
	data, err := readTasks()
	if err != nil {
		return fmt.Errorf("Error saving the Tasks file")
	}

	task.ID = len(data)

	data = append(data, task)

	strData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Error saving the Tasks file")
	}

	return os.WriteFile(fileName, strData, 0644)

}

func readTasks() ([]Task, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}, fmt.Errorf("Error reading the Tasks file")
	}

	var mapData []Task

	var errr error = json.Unmarshal(data, &mapData)
	if errr != nil {
		return []Task{}, fmt.Errorf("Error reading the Tasks file")
	}

	return mapData, nil
}

type Task struct {
	ID   int    `json:"id"`   // Capital I
	Task string `json:"task"` // Capital T
	Done bool   `json:"done"` // Capital D
}
