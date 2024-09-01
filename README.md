# Todo List Application

A command-line todo list application written in Go. Manage tasks with commands to add, list, complete, and remove tasks. Data is stored in a JSON file.

## Features

- Add new tasks
- List tasks (all, pending, completed)
- Mark tasks as completed
- Remove tasks
- Persistent storage with JSON

## Requirements

- Go 1.18+
- Windows, macOS, or Linux

## Setup

1. **Clone the Repository**
   ```bash
   git clone https://github.com/JakubSchwenkbeck/CLI-ToDoList.git
   cd CLI-ToDoLit
   ```
   ## Setup

2. **Run the Application**

   **On Windows:**

   Use the provided batch file `run_todo.bat` to simplify running the application:
   ```batch
   run_todo.bat [command] [arguments...]
   ```

 ###  Examples:

  ```batch
run_todo.bat add "Buy groceries"
run_todo.bat list all
run_todo.bat complete 1
run_todo.bat remove 1
   ```
