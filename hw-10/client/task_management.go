package client

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-10/jsonutil"
	"encoding/json"
	"fmt"
)

func AddTask(title string, description string) error {
	taskList, err := jsonutil.OpenJSONFile("tasks.json")
	if err != nil {
		return fmt.Errorf("Failed to open JSON file: %w", err)
	}

	var task Tasks
	err = json.Unmarshal(taskList, &task)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal task list: %w", err)
	}

	newTask := Task{
		Id:          len(task.Tasks) + 1,
		Title:       title,
		Description: description,
		Status:      false,
	}
	task.Tasks = append(task.Tasks, newTask)

	updatedTaskList, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated task list: %w", err)
	}

	err = jsonutil.SaveJSONFile("tasks.json", updatedTaskList)
	if err != nil {
		return fmt.Errorf("Failed to save updated task list: %w", err)
	}

	return nil
}

func UpdateTaskStatus(id int, newStatus bool) error {
	taskList, err := jsonutil.OpenJSONFile("tasks.json")
	if err != nil {
		return fmt.Errorf("Failed to open JSON file: %w", err)
	}

	var task Tasks
	err = json.Unmarshal(taskList, &task)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal task list: %w", err)
	}

	for i, val := range task.Tasks {
		if val.Id == id {
			task.Tasks[i].Status = newStatus
		}
	}

	updatedTaskList, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated task list: %w", err)
	}

	err = jsonutil.SaveJSONFile("tasks.json", updatedTaskList)
	if err != nil {
		return fmt.Errorf("Failed to save updated task list: %w", err)
	}

	return nil
}

func RemoveTask(id int) error {
	taskList, err := jsonutil.OpenJSONFile("tasks.json")
	if err != nil {
		return fmt.Errorf("Failed to open JSON file: %w", err)
	}

	var tasks Tasks
	err = json.Unmarshal(taskList, &tasks)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal task list: %w", err)
	}

	for i, task := range tasks.Tasks {
		if task.Id == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			break
		}
	}

	updatedTaskList, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("Failed to marshal updated task list: %w", err)
	}

	err = jsonutil.SaveJSONFile("tasks.json", updatedTaskList)
	if err != nil {
		return fmt.Errorf("Failed to save updated task list: %w", err)
	}

	return nil
}
