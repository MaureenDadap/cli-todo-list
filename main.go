package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

const TODO_LIST_FILE = "tasks.csv"
const TIME_FORMAT = time.RFC850

func main() {
	printMainHeader()

	list := List{}
	list.tasks = GetTaskFromFile()

	askForInput(&list)
	printTasksTable(list.tasks)

	list.Save()
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

func askForInput(list *List) {
	reader := bufio.NewReader(os.Stdin)

	done := false

	i := list.GetLastID() + 1
	for {
		if done {
			break
		}

		println("Enter a new task:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")

		currentTime := time.Now().Format(TIME_FORMAT)

		newTask := Task{i, input, currentTime, false}

		list.Add(newTask)

		done = askIfDone()

		i++
	}
}

func printTasksTable(tasks []Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Task", "Created At", "Status"})

	for _, task := range tasks {
		status := "Pending"

		if task.isDone {
			status = "Done"
		}

		t.AppendRow(table.Row{task.id, task.name, task.createdAt, status})
	}

	t.Render()
}
