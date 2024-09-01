package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [arguments...]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add <task>")
			return
		}
		task := strings.Join(os.Args[2:], " ")
		if err := addTask(task); err != nil {
			fmt.Println("Error adding task:", err)
		} else {
			fmt.Println("Task added:", task)
		}
	case "list":
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "all", "pending", "completed":
				if err := listTasks(os.Args[2]); err != nil {
					fmt.Println("Error listing tasks:", err)
				}
			default:
				fmt.Println("Usage: todo list <all|pending|completed>")
			}
		} else {
			if err := listTasks("all"); err != nil {
				fmt.Println("Error listing tasks:", err)
			}
		}
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo complete <task_index>")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task index:", err)
			return
		}
		if err := markTaskCompleted(index - 1); err != nil {
			fmt.Println("Error marking task completed:", err)
		} else {
			fmt.Println("Task marked as completed")
		}
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo remove <task_index>")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task index:", err)
			return
		}
		if err := removeTask(index - 1); err != nil {
			fmt.Println("Error removing task:", err)
		} else {
			fmt.Println("Task removed")
		}
	default:
		fmt.Println("Usage: todo <command> [arguments...]")
		fmt.Println("Commands: add, list, complete, remove")
	}
}
