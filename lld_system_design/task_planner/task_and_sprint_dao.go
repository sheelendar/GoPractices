package main

import (
	"fmt"
)

type Status int
type TaskType int

const (
	INIT = iota + 1
	COMPLATED
	DONE
	IN_PROGRESS
	IN_TESTING
	IN_QA

	STORY = iota + 1
	TASK
	BUG
)

var uniqueID = 1

func getUniqueID() int {
	id := uniqueID
	uniqueID++
	return id

}

type Task struct {
	id       int
	name     string
	subtract string
	user     User
	taskType TaskType
	status   Status
}

func TaksConstracter() Task {
	return Task{id: getUniqueID(), status: INIT}
}

func (task *Task) GetId() int {
	return task.id
}

func (task *Task) GetName() string {
	return task.name
}

func (task *Task) GetSubtract() string {
	return task.subtract
}

func (task *Task) GetUser() User {
	return task.user
}

func (task *Task) GetTaskType() TaskType {
	return task.taskType
}

func (task *Task) GetStatus() Status {
	return task.status
}

func (task *Task) SetName(name string) *Task {
	task.name = name
	return task
}

func (task *Task) SetSubtract(subtract string) *Task {
	task.subtract = subtract
	return task
}

func (task *Task) SetUser(user User) *Task {
	task.user = user
	return task
}

func (task *Task) SetTaskType(taskType TaskType) *Task {
	task.taskType = taskType
	return task
}

func (task *Task) SetStatus(status Status) *Task {
	task.status = status
	return task
}

// / ............................Sprint...........................
type Sprint struct {
	name  string
	begin int
	end   int
	tasks []Task
}

func SprintConstracter(name string, begin, end int) Sprint {
	return Sprint{name: name, begin: begin, end: end}
}

func (sprint *Sprint) addTask(task Task) {
	sprint.tasks = append(sprint.tasks, task)
}

func (sprint *Sprint) printSpring() {
	fmt.Println("Sprint name: ", sprint.name, " Begin: ", sprint.begin, " End: ", sprint.end)
	fmt.Println("task details")
	for _, task := range sprint.tasks {
		fmt.Println("task name ", task.name, " task status ", task.status)
	}
}

// .............setter and getter functions..................
func (sprint *Sprint) GetBegin() int {
	return sprint.begin
}

func (sprint *Sprint) GetEnd() int {
	return sprint.end
}

func (sprint *Sprint) GetTasks() []Task {
	return sprint.tasks
}

func (sprint *Sprint) SetBegin(begin int) *Sprint {
	sprint.begin = begin
	return sprint
}

func (sprint *Sprint) SetEnd(end int) *Sprint {
	sprint.end = end
	return sprint
}

func (sprint *Sprint) SetTasks(tasks []Task) *Sprint {
	sprint.tasks = tasks
	return sprint
}

func (sprint *Sprint) GetName() string {
	return sprint.name
}

func (sprint *Sprint) SetName(name string) *Sprint {
	sprint.name = name
	return sprint
}
