package main

import "fmt"

/*
 Sprint: A sprint is a short, time boxed period when a team works to complete a set amount of work.
2. Task: Any feature scheduled in the sprint is broken down into specific technical tasks. These tasks are
used by teams to decompose user stories or items to a more granular level.

The task planner should have the following functionalities:
• User should be able to create Task of type Story, Feature, Bug. Each can have their own status
• Stories can have further subtracts
• Should be able to change the status of any task
• User should be able to create any sprint. Should be able to add any task to sprint and remove from it.
• User should be able to print
{
	▪ Delayed task(s)
	▪ Sprint details
	▪ Tasks assigned to the user
}
*/

func main() {
	fmt.Println(".................Task planner app..............")
	fmt.Println()
	user1 := User{}
	user2 := User{}

	task1 := user1.createTaskByType(TASK)
	task11 := user1.createTaskByType(BUG)

	task2 := user2.createTaskByType(BUG)
	task22 := user2.createTask("this is a subtract")

	sprint1 := user1.createSprint("sprint1", 10, 12)
	sprint2 := user2.createSprint("sprint2", 20, 25)

	fmt.Println(user1.changesTaskStatus(task11, IN_PROGRESS)) // true

	fmt.Println(user1.addToSprint(sprint1, task1))       // true
	fmt.Println(user1.addToSprint(sprint1, task11))      // true
	fmt.Println(user1.addToSprint(sprint2, task1))       // flase
	fmt.Println(user1.removeFromSprint(sprint1, task11)) // true

	fmt.Println(user2.addToSprint(sprint1, task1))      // false
	fmt.Println(user2.removeFromSprint(sprint1, task2)) // false
	fmt.Println(user2.addToSprint(sprint2, task1))      // false

	fmt.Println(user2.addToSprint(sprint2, task2))  // true
	fmt.Println(user2.addToSprint(sprint2, task22)) // true

	sprint1.printSpring()
	sprint2.printSpring()
	fmt.Println()
	fmt.Println("user task detials")
	user1.printAllTask()
	user2.printAllTask()

}
