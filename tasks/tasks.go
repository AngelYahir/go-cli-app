package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func ListTask(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show")
		return
	}

	for _, task := range tasks {
		status := "[\033[31mx\033[0m]"
		if task.Completed {
			status = "[\033[32m✓\033[0m]"
		}
		fmt.Printf("%s \033[35m%d\033[0m %s\n", status, task.ID, task.Name)
	}
}

func AddTask(tasks []Task, name string) []Task {
	task := Task{
		ID:        GetID(tasks),
		Name:      name,
		Completed: false,
	}

	tasks = append(tasks, task)

	return tasks
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}

	return tasks
}

func SaveTasks(file *os.File, tasks []Task) {
	//? Save the tasks to the file
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)

	if err != nil {
		panic(err)
	}

	err = writer.Flush()

	if err != nil {
		panic(err)
	}
}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			task.Completed = true
			tasks[i] = task
			break
		}
	}

	return tasks
}

func GetID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}
