package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task struct represents a single task
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// TodoList struct represents the todo list
type TodoList struct {
	Tasks []Task
}

// PrintTasks prints all tasks in the todo list
func (tl *TodoList) ViewTasks() {
	fmt.Println("Todo List:")
	for _, task := range tl.Tasks {
		completedStatus := "Incomplete"
		if task.Completed {
			completedStatus = "Completed"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, completedStatus)
	}
	fmt.Println()
}

// AddTask adds a new task to the todo list
func (tl *TodoList) AddTask(title string) {
	id := len(tl.Tasks) + 1
	task := Task{ID: id, Title: title, Completed: false}
	tl.Tasks = append(tl.Tasks, task)
	fmt.Println("Task added successfully!")
	fmt.Println()
}

// UpdateTask updates the status of a task in the todo list
func (tl *TodoList) UpdateTask(taskID int, completed bool) {
	for i, task := range tl.Tasks {
		if task.ID == taskID {
			tl.Tasks[i].Completed = completed
			fmt.Println("Task updated successfully!")
			fmt.Println()
			return
		}
	}
	fmt.Println("Task not found!")
	fmt.Println()
}

// DeleteTask deletes a task from the todo list
func (tl *TodoList) DeleteTask(taskID int) {
	for i, task := range tl.Tasks {
		if task.ID == taskID {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			fmt.Println()
			return
		}
	}
	fmt.Println("Task not found!")
	fmt.Println()
}

func main() {
	todoList := TodoList{}

	for {
		fmt.Println("Please select an option:")
		fmt.Println("1. View tasks")
		fmt.Println("2. Add a task")
		fmt.Println("3. Update a task")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Exit")

		var choice string
		fmt.Print("Enter Your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			todoList.ViewTasks()
		case "2":
			fmt.Print("Enter task title: ")
			var title string
			fmt.Scanln(&title)
			todoList.AddTask(title)
		case "3":
			fmt.Print("Enter task ID: ")
			var taskID int
			fmt.Scanln(&taskID)
			fmt.Print("Enter completion status (true/false): ")
			var completedInput string
			fmt.Scanln(&completedInput)
			completed, err := strconv.ParseBool(strings.TrimSpace(completedInput))
			if err != nil {
				fmt.Println("Invalid completion status!")
				fmt.Println()
				continue
			}
			todoList.UpdateTask(taskID, completed)
		case "4":
			fmt.Print("Enter task ID: ")
			var taskID int
			fmt.Scanln(&taskID)
			todoList.DeleteTask(taskID)
		case "5":
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice!")
			fmt.Println()
		}
	}
}
