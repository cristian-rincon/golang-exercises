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
		if task.done == done {
			fmt.Println("id", idx, ":", task.name, "|", task.description, "|", task.done)
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
		description: "Task description 1",
		done:        false,
	}
	t1 := &task{
		name:        "Task 2",
		description: "Task description 2",
		done:        true,
	}
	t2 := &task{
		name:        "Task 3",
		description: "Task description 3",
		done:        false,
	}
	t3 := &task{
		name:        "Task 4",
		description: "Task description 5",
		done:        false,
	}

	taskL := &taskList{
		tasks: []*task{t, t1, t2},
	}
	// fmt.Printf("%+v\n", taskL.tasks[0])
	taskL.addTask(t3)
	t3.markAsDone()
	taskL.printList(false)

	mapTask := make(map[string]*task)
	mapTask["1"] = t

	fmt.Println(mapTask["1"].name)

}
