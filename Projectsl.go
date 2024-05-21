package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var taskid = 0

// Define Task struct
type Task struct {
	id       int
	toDo     string
	isDone   bool
	date     string
	deadline string
}

// Function to add a task
func add(tasks []Task) []Task {
	fmt.Println("Please Enter the task")

	var localTask Task
	inp := receiver()
	fmt.Println(inp)
	taskid++
	localTask.id = taskid
	localTask.toDo = inp
	localTask.date = time.Now().Format("2006-01-02")
	tasks = append(tasks, localTask)
	return tasks
}

// Function to list all tasks
func listTasks(tasks []Task) {
	fmt.Println("---------- List of Tasks ---------")
	for _, value := range tasks {
		fmt.Printf("id: %d\t task:%s\t isDone:%t Date:%s\n", value.id, value.toDo, value.isDone, value.date)
	}
}

// Function to mark a task as done
func markAsDone(tasks []Task) []Task {
	fmt.Println("---------- Mark Task as Done ---------")
	fmt.Println("Enter the ID")
	input := receiver()
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Enter a valid Id")
		return tasks
	}
	for i, task := range tasks {
		if task.id == id {
			tasks[i].isDone = true
			break
		}
	}
	return tasks
}

// Function to delete a task
func deleteTask(tasks []Task) []Task {
	fmt.Println("---------- Delete Task ---------")
	fmt.Println("Enter the ID")
	input := receiver()
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Enter a valid Id")
		return tasks
	}
	for i, task := range tasks {
		if task.id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	return tasks
}

// Function to receive user input
func receiver() string {
	reader := bufio.NewReader(os.Stdin)
	changed, _ := reader.ReadString('\n')
	return changed
}

// Main function
func main() {
	var todoTasks []Task

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		input := receiver()
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		switch choice {
		case 1:
			todoTasks = add(todoTasks)
		case 2:
			listTasks(todoTasks)
		case 3:
			todoTasks = markAsDone(todoTasks)
		case 4:
			todoTasks = deleteTask(todoTasks)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
