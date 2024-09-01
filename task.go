package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

const filePath = "data.json"

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func loadTasks() ([]Task, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}
	return nil
}

func addTask(description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	task := Task{Description: description, Completed: false}
	tasks = append(tasks, task)
	return saveTasks(tasks)
}

func listTasks(filter string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		switch filter {
		case "all":
			fmt.Printf("%d. %s [%s]\n", index+1, task.Description, status(task.Completed))
		case "pending":
			if !task.Completed {
				fmt.Printf("%d. %s [%s]\n", index+1, task.Description, status(task.Completed))
			}
		case "completed":
			if task.Completed {
				fmt.Printf("%d. %s [%s]\n", index+1, task.Description, status(task.Completed))
			}
		}
	}
	return nil
}

func markTaskCompleted(index int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(tasks) {
		return errors.New("task index out of range")
	}

	tasks[index].Completed = true
	return saveTasks(tasks)
}

func removeTask(index int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(tasks) {
		return errors.New("task index out of range")
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	return saveTasks(tasks)
}

func status(completed bool) string {
	if completed {
		return "X"
	}
	return " "
}
