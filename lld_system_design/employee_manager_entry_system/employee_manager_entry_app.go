package main

import "fmt"

/*
Application Requirements
Create an employee structure in which employee has id, name and manager. Following functionalities are desired:
▪ Return employee details for a given employee id
▪ List all the subordinates of the given employee for a given id/name
▪ Given a name, search with the prefix search property.
The application should exist in memory
*/

func main() {
	fmt.Println("sheelendar")
	employee1 := getEmployeeConstractor("Sheelendar")
	employee2 := getEmployeeConstractor("Ram")
	employee3 := getEmployeeConstractor("Sayam")

	system := getSystemConstractor()
	system.registerEmployee(employee1)
	system.registerEmployee(employee2)
	system.registerEmployee(employee3)

	system.printDetails(employee2.id)
	fmt.Println("******************************************")

	system.printAllPrefixSearchName("S")
	fmt.Println("******************************************")

	system.registerManager(employee2.id, employee1.id)
	system.registerManager(employee3.id, employee1.id)

	system.printDetails(employee1.id)
	fmt.Println("******************************************")
}
