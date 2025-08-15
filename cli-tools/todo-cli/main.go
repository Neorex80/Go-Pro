package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	ID     int
	Task   string
	Done   bool
}

var todos []Todo
var nextID = 1

func main() {
	// Load existing todos from file
	loadTodos()

	// Command-line interface
	for {
		fmt.Println("\n===== Todo CLI =====")
		fmt.Println("1. List todos")
		fmt.Println("2. Add todo")
		fmt.Println("3. Complete todo")
		fmt.Println("4. Delete todo")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			listTodos()
		case "2":
			fmt.Print("Enter task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			addTodo(task)
		case "3":
			fmt.Print("Enter todo ID to complete: ")
			var id int
			fmt.Scanf("%d", &id)
			completeTodo(id)
		case "4":
			fmt.Print("Enter todo ID to delete: ")
			var id int
			fmt.Scanf("%d", &id)
			deleteTodo(id)
		case "5":
			// Save todos to file before exiting
			saveTodos()
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	fmt.Println("\nTodos:")
	for _, todo := range todos {
		status := " "
		if todo.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", todo.ID, status, todo.Task)
	}
}

func addTodo(task string) {
	if task == "" {
		fmt.Println("Task cannot be empty.")
		return
	}

	todo := Todo{
		ID:   nextID,
		Task: task,
		Done: false,
	}

	todos = append(todos, todo)
	nextID++
	fmt.Printf("Added todo: %s\n", task)
}

func completeTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = true
			fmt.Printf("Completed todo: %s\n", todo.Task)
			return
		}
	}
	fmt.Printf("Todo with ID %d not found.\n", id)
}

func deleteTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Printf("Deleted todo: %s\n", todo.Task)
			return
		}
	}
	fmt.Printf("Todo with ID %d not found.\n", id)
}

func saveTodos() {
	file, err := os.Create("todos.txt")
	if err != nil {
		fmt.Printf("Error saving todos: %v\n", err)
		return
	}
	defer file.Close()

	for _, todo := range todos {
		status := "0"
		if todo.Done {
			status = "1"
		}
		line := fmt.Sprintf("%d|%s|%s\n", todo.ID, status, todo.Task)
		file.WriteString(line)
	}

	// Update nextID to be one more than the highest ID
	if len(todos) > 0 {
		maxID := 0
		for _, todo := range todos {
			if todo.ID > maxID {
				maxID = todo.ID
			}
		}
		nextID = maxID + 1
	}
}

func loadTodos() {
	file, err := os.Open("todos.txt")
	if err != nil {
		// File doesn't exist, which is fine for first run
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 3 {
			var todo Todo
			fmt.Sscanf(parts[0], "%d", &todo.ID)
			if parts[1] == "1" {
				todo.Done = true
			}
			todo.Task = parts[2]
			todos = append(todos, todo)

			// Update nextID to be one more than the highest ID
			if todo.ID >= nextID {
				nextID = todo.ID + 1
			}
		}
	}
}
