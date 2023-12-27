package main

import "fmt"

type User struct {
	taskList   []*Task
	sprintlist []*Sprint
}

func (user *User) createTaskByType(taskType TaskType) *Task {
	if taskType == STORY {
		fmt.Println("Task of type Story is being created with no subtract")
	}
	task := TaksConstracter()
	task.SetTaskType(taskType)
	task.SetUser(*user)
	user.taskList = append(user.taskList, &task)
	return &task
}

func (user *User) createTask(s string) *Task {

	task := TaksConstracter()
	task.SetTaskType(STORY)
	task.SetSubtract(s)
	task.SetUser(*user)
	user.taskList = append(user.taskList, &task)
	return &task
}

func (user *User) createSprint(s string, begin, end int) *Sprint {
	sprint := Sprint{name: s, begin: begin, end: end}
	user.sprintlist = append(user.sprintlist, &sprint)
	return &sprint

}

func (user *User) addToSprint(sprint *Sprint, task *Task) bool {
	var taskToBeAdd *Task
	for _, t := range user.taskList {
		if t.GetId() == task.GetId() {
			taskToBeAdd = task
			break
		}
	}
	if taskToBeAdd == nil {
		return false
	}

	for _, spt := range user.sprintlist {
		if spt.GetName() == sprint.GetName() {
			sprint.addTask(*taskToBeAdd)
			return true
		}
	}
	return false
}

func (user *User) removeFromSprint(sprint *Sprint, task *Task) bool {

	var userSprintFind *Sprint
	for _, spt := range user.sprintlist {
		if spt.GetName() == sprint.GetName() {
			userSprintFind = sprint
			break
		}
	}
	if userSprintFind == nil {
		return false
	}
	var taskList []Task
	for _, t := range userSprintFind.GetTasks() {
		if t.id != task.id {
			taskList = append(taskList, t)
		}
	}
	sprint.SetTasks(taskList)
	return true
}

func (user *User) printAllTask() {
	for _, task := range user.taskList {
		fmt.Println("task name ", task.id, " task status ", task.status)
	}

}

func (user *User) changesTaskStatus(task *Task, status Status) bool {
	for _, t := range user.taskList {
		if t.GetId() == task.GetId() {
			task.SetStatus(status)
			return true
		}
	}
	return false
}
