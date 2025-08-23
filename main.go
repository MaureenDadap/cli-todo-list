package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	printMainHeader()

	tasks := askForInput()

	printMainHeader()

	printTasksTable(tasks)
}

type Task struct {
	id         int
	name       string
	created_at string
	is_done    bool
}

func printMainHeader() {
	println("====== My TODO List ======= ")
}

func askIfDone() bool {
	print("Enter another one? (y/n): ")

	input := ""

	fmt.Scan(&input)

	if input != "y" && input != "n" {
		println("Invalid Input!")
		askIfDone()
	}

	if input == "y" {
		return false
	}

	return true
}

func askForInput() []Task {
	tasks := []Task{}
	reader := bufio.NewReader(os.Stdin)

	done := false

	i := 1
	for {
		if done {
			break
		}

		println("Enter a new task:")
		input, _ := reader.ReadString('\n')

		currentTime := time.Now().Format(time.RFC850)

		newTask := Task{i, input, currentTime, false}

		tasks = append(tasks, newTask)

		done = askIfDone()

		i++
	}

	return tasks
}

func printTasksTable(tasks []Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Task", "Created At", "Status"})

	for _, task := range tasks {
		status := "Pending"

		if task.is_done {
			status = "Done"
		}

		t.AppendRow(table.Row{
			task.id,
			task.name,
			task.created_at,
			status,
		})
	}

	t.Render()
}
