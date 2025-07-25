package main

import "fmt"

type Task struct {
	Title       string
	Completed   bool
	Priority    int
	Description string
}

func main() {
	task1 := Task{
		Title:       "Learn GoLang",
		Completed:   false,
		Priority:    1,
		Description: "get used to the syntax and features of golang",
	}

	task2 := Task{
		Title:       "Push code to repo",
		Completed:   true,
		Priority:    2,
		Description: "archive work history in github",
	}

	tasks := []Task{task1, task2}

	for _, v := range tasks {
		fmt.Print(stringTodo(v))
	}
}

func stringTodo(t Task) string {
	msg := fmt.Sprintf("\n[ %s ]\n %s \n Priority: %d\n", t.Title, t.Description, t.Priority)
	return msg
}
