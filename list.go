package main

import "strconv"

type List struct {
	tasks []Task
}

type Task struct {
	id        int
	name      string
	createdAt string
	isDone    bool
}

func (t *Task) idString() string {
	return strconv.Itoa(t.id)
}

func (t *Task) isDoneString() string {
	if t.isDone {
		return "1"
	}

	return "0"
}

func GetTaskFromFile() []Task {
	data := ReadAndParseCSV(TODO_LIST_FILE)

	if data.GetTotalRows() < 1 {
		return nil
	}

	tasks := []Task{}

	for _, row := range data.rows {
		isDone := row.columns[3] == "1"

		task := Task{
			id:        Must(strconv.Atoi(row.columns[0])),
			name:      row.columns[1],
			createdAt: row.columns[2],
			isDone:    isDone,
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func (l *List) Add(task Task) {
	l.tasks = append(l.tasks, task)
}

func (l *List) GetLastID() int {
	i := len(l.tasks)

	if i < 1 {
		return 0
	}

	lastTask := l.tasks[i-1]

	return lastTask.id
}

// Save list to storage file
func (l *List) Save() {
	csvHeaders := getHeaders()

	csv := CSV{}
	csv.SetHeader(ToRow(csvHeaders))

	for _, task := range l.tasks {
		data := []string{
			task.idString(),
			task.name,
			task.createdAt,
			task.isDoneString()}

		csv.AppendRow(ToRow(data))
	}

	csv.SaveToFile(TODO_LIST_FILE)
}
