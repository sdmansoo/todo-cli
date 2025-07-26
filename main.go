package main

import "fmt"

type Task struct {
	Title       string
	Description string
	Completed   bool
	Priority    int
}

func main() {
	task1 := Task{
		Title:       "Learn GoLang",
		Description: "get used to the syntax and features of golang",
		Completed:   false,
		Priority:    1,
	}

	task2 := Task{
		Title:       "Push code to repo",
		Description: "archive work history in github",
		Completed:   true,
		Priority:    2,
	}

	task3 := addTask("CRUD", "add create,update,read,delete functions", 2)

	task4 := addTask("Fake", "delete this task", 3)

	tasks := []Task{task1, task2, task3, task4}

	printTaskList(tasks)

	tasks = deleteTask(tasks, 4)
	fmt.Println("\ndeleting task 4\n ")

	printTaskList(tasks)

	fmt.Println("\nMarking task 3 as complete\n ")
	tasks[2] = markComplete(tasks[2])

	printTaskList(tasks)
}

func stringTask(t Task) string {
	var status string
	if t.Completed {
		status = "✅"
	} else {
		status = "❌"
	}

	msg := fmt.Sprintf("[ %s ] %s\n %s \n Priority: %d\n",
		t.Title, status, t.Description, t.Priority)
	return msg
}

func printTaskList(taskList []Task) {
	for i, v := range taskList {
		task := stringTask(v)
		fmt.Printf("%d %s", i+1, task)
	}
}

func addTask(title string, description string, priority int) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,
		Priority:    priority,
	}
}

func markComplete(t Task) Task {
	return Task{
		Title:       t.Title,
		Description: t.Description,
		Completed:   true,
		Priority:    t.Priority,
	}
}

func deleteTask(taskList []Task, index int) []Task {
	return append(taskList[:index-1], taskList[index:]...)
}
