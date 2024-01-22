package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// const (
// 	todo = iota
// 	in_progress
// 	done
// 	archived
// )

// var Items map[int]Item

type Item struct {
	ID   int
	Name string
	// Description string
	// Status int
}

var todoList []Item
var lastID int

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task Description: ")
	scanner.Scan()
	text := scanner.Text()
	lastID++
	item := Item{ID: lastID, Name: text}
	todoList = append(todoList, item)
	fmt.Printf("Task Added: %s\n", text)
}

func listTasks() {
	if len(todoList) == 0 {
		fmt.Println("Todo List empty")
		return
	}
	fmt.Println("todo List: ")
	for _, task := range todoList {
		fmt.Printf("%d. %s\n", task.ID, task.Name)
	}
}

func removeTask(scanner *bufio.Scanner) {
	if len(todoList) == 0 {
		fmt.Println("No tasks to remove")
	}

	fmt.Printf("Enter task ID to remove: ")
	scanner.Scan()
	taskID, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Enter valid number")
		return
	}
	idx := -1
	for index, task := range todoList {
		if task.ID == taskID {
			idx = index
			break
		}
	}
	if idx == -1 {
		fmt.Println("task not found")
		return
	}
	todoList = append(todoList[:idx], todoList[idx+1:]...)
	fmt.Printf("task with id %d removed\n", taskID)
}
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. List Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Exit")
		fmt.Print("Enter you choice: ")

		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid choice. Enter number")
			continue
		}
		switch choice {
		case 1:
			listTasks()
		case 2:
			addTask(scanner)
		case 3:
			removeTask(scanner)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter valid option")
		}

	}
	// item1 := Item{1,"test-1","test description",done}
	// Items = make(map[int]Item)
}
