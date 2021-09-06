package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) addTask(task *task) {
	t.tasks = append(t.tasks, task)
}

func (t *taskList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printList(done bool) {
	for idx, task := range t.tasks {
		if task.done {
			fmt.Println("Index", idx, ":", task.name, "is done")
		}	else {
			fmt.Println("Index", idx, ":", task.name, "is not done")
		}
	}
}

type task struct {
	name        string
	description string
	done        bool
}

func (t *task) markAsDone() string {
	t.done = true
	return t.name
}

func (t *task) updateDescription(newDescription string) string {
	t.description = newDescription
	return t.name
}

func (t *task) updateName(newName string) string {
	t.name = newName
	return t.name
}

func main() {
	t := &task{
		name:        "Task 1",
		description: "Task description",
		done:        false,
	}
	t1 := &task{
		name:        "Task 2",
		description: "Task description",
		done:        true,
	}
	t2 := &task{
		name:        "Task 3",
		description: "Task description",
		done:        false,
	}
	t3 := &task{
		name:        "Task 4",
		description: "Task description",
		done:        false,
	}

	taskL := &taskList{
		tasks: []*task{t, t1, t2},
	}
	// fmt.Printf("%+v\n", taskL.tasks[0])
	taskL.addTask(t3)
	t3.markAsDone()
	taskL.printList(true)
}
