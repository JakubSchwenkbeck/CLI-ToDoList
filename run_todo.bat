@echo off
REM Compile and run the Go application

REM Check if Go is installed
where go >nul 2>&1
if %errorlevel% neq 0 (
    echo Go is not installed. Please install Go and add it to your PATH.
    exit /b 1
)

REM Run the Go application with the provided arguments
go run todo_list.go task.go %*
