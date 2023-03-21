package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/AngelYahir/go-cli-app/tasks"
)

func main() {

	file, err := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)

		if err != nil {
			panic(err)
		}

		//? Convert bytes to json and assign to tasks
		err = json.Unmarshal(bytes, &tasks)

		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	//? Get the command line arguments
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTask(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Task name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		fmt.Println("\033[32mTask added successfully\033[0m")
		task.SaveTasks(file, tasks)

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("\033[33mPlease provide the task ID\033[0m")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("\033[31mInvalid ID\033[0m")
			return
		}

		tasks = task.DeleteTask(tasks, id)

		task.SaveTasks(file, tasks)

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("\033[33mPlease provide the task ID\033[0m")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("\033[31mInvalid ID\033[0m")
			return
		}

		tasks = task.CompleteTask(tasks, id)

		task.SaveTasks(file, tasks)

	}

}

func printUsage() {
	fmt.Println("\033[33mUsage: Go CLI App [\033[34mList \033[0m| \033[36mAdd\033[0m | \033[31mRemove\033[0m | \033[32mComplete\033[33m]\033[0m")
}
