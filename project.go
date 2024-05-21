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

//create a structure with parameters of task

type Task struct {
	id       int
	toDo     string
	isDone   bool
	date     string
	deadline string
}

func add(t map[int]Task) {
	fmt.Println("Please Enter the task")
	
	var localTask Task
	inp := receiver()
	fmt.Println(inp)
	localTask.id = taskid + 1
	localTask.toDo = inp
	localTask.date = time.Now().Format("2006-01-02")
    t[taskid] = localTask



}
func listTask(t map[int]Task)  {
    fmt.Println("---------- List of Tasks ---------")
    for _,value := range(t){
        fmt.Println()
        fmt.Printf("id: %d\t task:%s\t isDone:%t Date:%s", value.id, value.toDo, value.isDone, value.date)
        fmt.Println()
    }
 
}

func receiver() string {
	reader := bufio.NewReader(os.Stdin)
	changed, _ := reader.ReadString('\n')
	return changed

}
func isDone(t map[int]Task) {
    fmt.Println("---------- Is Done  ---------")
	fmt.Println("Enter the ID")
    input := receiver()
    id, err := strconv.Atoi(strings.TrimSpace(input))
    if err != nil{
        fmt.Println("Enter a valid Id")
    } else{
        lt := t[id]
        lt.isDone = true
        t[id] = lt
    }
   
 
}
 
func deleteTask(t map[int]Task) {
    input := receiver()
    id, err := strconv.Atoi(strings.TrimSpace(input))
    if err != nil{
        fmt.Println("Enter a valid Id")
    } else{
        delete(t, id)
    }
 
}

// Main function to display the menu and handle user input
func main() {

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
		fmt.Println(choice)
		todoTask := make(map[int]Task)

		switch choice {
			case 1: add(todoTasks)
			case 2: listTask(todoTasks)
			case 3: isDone(todoTasks)
			case 4: deleteTask(todoTasks)
			case 5: os.Exit(0)
			default:
		}
	}
}
