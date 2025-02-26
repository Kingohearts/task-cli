package main

import (
	"fmt"
	"os"
	"strconv"
)

const sourceFile = "tasks.json"

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func initFile() {
	var file *os.File
	var err error
	if !fileExists(sourceFile) {
		file, err = os.Create(sourceFile)
		if err != nil {
			fmt.Printf("Error Creating file: %v", err)
		}
		defer file.Close()
	}
}

func main() {
	initFile()
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage task-cli <command> [args] ")
		return
	}

	command := args[1]
	switch command {
	case "add":
		if len(args) == 3 {
			AddTask(sourceFile, args[2])
			fmt.Println("Task added succesfully")
		} else {
			fmt.Println("Invalid Command")
			return
		}
	case "update":
		if len(args) == 4 {
			id, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("Error Occured while converting string to integer %v", err)
			}
			UpdateDescription(sourceFile, id, args[3])
			fmt.Println("Task updated succesfully")
		} else {
			fmt.Println("Invalid Command")
			return
		}
	case "mark-in-progress":
		if len(args) == 3 {
			id, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("Error Occured while converting string to integer %v", err)
			}
			UpdateStatus(sourceFile, id, "in-progress")
			fmt.Println("Task updated succesfully")
		} else {
			fmt.Println("Invalid Command")
			return
		}
	case "mark-done":
		if len(args) == 3 {
			id, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("Error Occured while converting string to integer %v", err)
			}
			UpdateStatus(sourceFile, id, "done")
			fmt.Println("Task updated succesfully")
		} else {
			fmt.Println("Invalid Command")
			return
		}
	case "delete":
		if len(args) == 3 {
			id, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("Error Occured while converting string to integer %v", err)
			}
			DeleteTask(sourceFile, id)
			fmt.Println("Task deleted succesfully")
		} else {
			fmt.Println("Invalid Command")
			return
		}
	case "list":
		if len(args) == 2 {
			tasks, err := GetAllTasks(sourceFile)
			if err != nil {
				fmt.Printf("Error Occured while getting tasks %v", err)
			}
			fmt.Printf("Tasks are %v", tasks)
		} else if len(args) == 3 {
			tasks, err := GetTasksByStatus(sourceFile, args[2])
			if err != nil {
				fmt.Printf("Error Occured while getting tasks %v", err)
			}
			fmt.Printf("Tasks are %v", tasks)
		} else {
			fmt.Println("Invalid Command")
			return
		}
	default:
		return
	}
}
