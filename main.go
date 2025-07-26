package main

import "fmt"

type Task struct {
	Title       string
	Description string
	Completed   bool
	Priority    int
}

type TaskList struct {
	tasks []Task
}

func main() {
	var taskList = TaskList{}

	taskList.addTask("Learn GoLang", "get used to the syntax and features of golang", 1)

	taskList.addTask("Push code to repo", "archive work history in github", 2)

	taskList.addTask("CRUD", "add create,update,read,delete functions", 2)

	taskList.addTask("Fake", "delete this task", 3)

	taskList.printTaskList()

	taskList.deleteTask(4)
	fmt.Println("\ndeleting task 4\n ")

	taskList.printTaskList()

	fmt.Println("\nMarking task 3 as complete\n ")
	taskList.markComplete(3)

	taskList.printTaskList()
}

func (t *Task) stringTask() string {
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

func (l *TaskList) printTaskList() {
	for i, v := range l.tasks {
		task := v.stringTask()
		fmt.Printf("%d %s", i+1, task)
	}
}

func (l *TaskList) addTask(title string, description string, priority int) {
	t := Task{title, description, false, priority}
	l.tasks = append(l.tasks, t)
}

func (l *TaskList) markComplete(taskNumber int) {
	l.tasks[taskNumber-1].Completed = true
}

func (l *TaskList) deleteTask(taskNumber int) {
	l.tasks = append(l.tasks[:taskNumber-1], l.tasks[taskNumber:]...)
}
