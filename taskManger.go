package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

const timeFormat = "2006-01-02 15:04:05"

func AddTask(fileName string, description string) error {

	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}
	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err == io.EOF {
		tasks = []Task{} // Ensures proper JSON encoding
	} else if err != nil {
		return fmt.Errorf("error %v", err)
	}

	var id int
	var totalTasks = len(tasks)
	if totalTasks == 0 {
		id = 1
	} else {
		id = tasks[totalTasks-1].Id + 1
	}
	tasks = append(tasks, Task{
		Id:          id,
		Description: description,
		Status:      "to-do",
		CreatedAt:   time.Now().Format(timeFormat),
		UpdatedAt:   "",
	})
	return reWriteJSON(file, tasks)
}

func UpdateDescription(fileName string, id int, description string) error {

	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}
	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err == io.EOF {
		tasks = []Task{} // Ensures proper JSON encoding
	} else if err != nil {
		return fmt.Errorf("error %v", err)
	}

	index, err := findTaskIndex(id, tasks)
	if err != nil {
		return fmt.Errorf("error  reason %v", err)
	}
	tasks[index].Description = description
	tasks[index].UpdatedAt = time.Now().Format(timeFormat)

	return reWriteJSON(file, tasks)
}

func UpdateStatus(fileName string, id int, status string) error {

	if status != "in-progress" && status != "done" {
		return fmt.Errorf("error invalid status value")
	}

	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}
	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err == io.EOF {
		tasks = []Task{} // Ensures proper JSON encoding
	} else if err != nil {
		return fmt.Errorf("error %v", err)
	}

	index, err := findTaskIndex(id, tasks)
	if err != nil {
		return fmt.Errorf("error  reason %v", err)
	}
	tasks[index].Status = status
	tasks[index].UpdatedAt = time.Now().Format(timeFormat)

	return reWriteJSON(file, tasks)
}

func DeleteTask(fileName string, id int) error {

	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}
	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err == io.EOF {
		tasks = []Task{} // Ensures proper JSON encoding
	} else if err != nil {
		return fmt.Errorf("error %v", err)
	}

	index, err := findTaskIndex(id, tasks)
	if err != nil {
		return fmt.Errorf("error  reason %v", err)
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	reWriteJSON(file, tasks)
	return nil
}

func GetTasksByStatus(fileName string, status string) ([]Task, error) {
	if status != "to-do" && status != "in-progress" && status != "done" {
		return nil, fmt.Errorf("error invalid status value")
	}
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("error %v", err)
	}
	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err == io.EOF {
		tasks = []Task{} // Ensures proper JSON encoding
	} else if err != nil {
		return nil, fmt.Errorf("error %v", err)
	}

	filteredTasks := []Task{}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].Status == status {
			filteredTasks = append(filteredTasks, tasks[i])
		}
	}

	return filteredTasks, nil
}

func GetAllTasks(fileName string) ([]Task, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("error %v", err)
	}
	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err == io.EOF {
		tasks = []Task{} // Ensures proper JSON encoding
	} else if err != nil {
		return nil, fmt.Errorf("error %v", err)
	}

	return tasks, nil
}

func findTaskIndex(id int, t []Task) (int, error) {
	var index int
	for i := 0; i < len(t); i++ {
		if t[i].Id == id {
			index = i
			return index, nil
		}
	}
	return -1, fmt.Errorf("error no matching Id")
}

func reWriteJSON(f *os.File, t []Task) error {
	f.Seek(0, 0)
	f.Truncate(0)
	err := json.NewEncoder(f).Encode(t)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}
	return nil

}
