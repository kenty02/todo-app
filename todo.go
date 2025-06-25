package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}
	task := Task{
		ID:    nextID(tasks),
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, task)
	err = saveTasks(tasks)
	 if err != nil {
		panic(err)
	}
}

func ListTasks() {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}
	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
	}
}


// 初見 とりあえず書いてみた
/*
func CompleteTask(id int) {
	task, err := loadTasks()
	if err != nil{
		panic(err)
	}
	task{
		ID: nextID(task),
		Title: title,
		Done: true,
	}
	
	if err := saveTasks(task); err != nil{
		panic(err)
	}
}
*/

func CompleteTask(id int){
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id{
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found{
		println("指定されたIDのタスクが見つかりませんでした｡")
		return
	}

	if err := saveTasks(tasks); err != nil {
		panic(err)
	}
}

func DeleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}

	for i, task := range tasks{
		if task.ID == id{
			if err := saveTasks(append(tasks[:i], tasks[i+1:]...)); err != nil {
				panic(err)
			}
			return
		}
	} 
}

