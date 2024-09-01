package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var filePath = "todo.txt"

// Function to add a task
func addTask(task string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(task + "\n"); err != nil {
		return err
	}
	return nil
}

// Function to list all tasks
func listTasks() error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

// Main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo_list <command> [task]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo_list add <task>")
			return
		}
		task := strings.Join(os.Args[2:], " ")
		if err := addTask(task); err != nil {
			fmt.Println("Error adding task:", err)
		} else {
			fmt.Println("Task added:", task)
		}
	case "list":
		if err := listTasks(); err != nil {
			fmt.Println("Error listing tasks:", err)
		}
	default:
		fmt.Println("Usage: todo_list <command> [task]")
		fmt.Println("Commands: add, list")
	}
}
