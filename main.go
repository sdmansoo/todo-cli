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

	for i, v := range tasks {
		fmt.Printf("%d ", i+1)
		fmt.Print(stringTask(v))
	}

	tasks = deleteTask(tasks, 4)
	fmt.Println("\ndeleting task 4\n ")

	for i, v := range tasks {
		fmt.Printf("%d ", i+1)
		fmt.Print(stringTask(v))
	}
}

func stringTask(t Task) string {
	msg := fmt.Sprintf("\n[ %s ]\n %s \n Priority: %d\n", t.Title, t.Description, t.Priority)
	return msg
}

func addTask(title string, description string, priority int) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,
		Priority:    priority,
	}
}

func deleteTask(taskList []Task, index int) []Task {
	return append(taskList[:index-1], taskList[index:]...)
}
