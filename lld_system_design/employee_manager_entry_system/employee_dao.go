package main

import (
	"fmt"
	"strings"
)

var uniqueID = 1
var uniqueManagerID = 1

func getUniqueID() int {
	id := uniqueID
	uniqueID++
	return id
}

func getManagerUniqueID() int {
	id := uniqueManagerID
	uniqueManagerID++
	return id
}

type Employee struct {
	id           int
	name         string
	subordinates []*Employee
	managerID    int
}

func getEmployeeConstractor(name string) *Employee {
	return &Employee{id: getUniqueID(), name: name}
}

func getEmployeeAndManagerConstractor(name string) *Employee {
	return &Employee{id: getUniqueID(), name: name, managerID: getManagerUniqueID()}
}

func (emp *Employee) addEmployee(employee *Employee) {
	emp.subordinates = append(emp.subordinates, employee)
}

func (employee *Employee) GetId() int {
	return employee.id
}

func (employee *Employee) GetName() string {
	return employee.name
}

func (employee *Employee) GetSubordinates() []*Employee {
	return employee.subordinates
}

func (employee *Employee) GetManagerID() int {
	return employee.managerID
}
func (employee *Employee) setManagerID(managerID int) {
	employee.managerID = managerID
}

// ****************************************** System *************************************
type System struct {
	listEmployee []*Employee
	employeeMap  map[int]*Employee
}

func getSystemConstractor() *System {
	return &System{employeeMap: make(map[int]*Employee)}
}

func (system *System) registerEmployee(employee *Employee) {
	system.listEmployee = append(system.listEmployee, employee)
	system.employeeMap[employee.GetId()] = employee
}

func (system *System) registerManager(empID, managerID int) {
	if system.employeeMap[empID] == nil || system.employeeMap[managerID] == nil {
		fmt.Println("Either Employee or manager not registerd. Please provide correct empId or managerID")
		return
	}
	system.employeeMap[empID].setManagerID(managerID)
	system.employeeMap[managerID].addEmployee(system.employeeMap[empID])
}
func (system *System) printDetails(empID int) {
	if system.employeeMap[empID] == nil {
		fmt.Println("employeeId not found into system")
		return
	}
	emp := system.employeeMap[empID]
	fmt.Println("id: ", emp.GetId(), "name: ", emp.GetName(), "managerID: ", emp.GetManagerID())
	if emp.GetSubordinates() == nil {
		return
	}
	fmt.Println("suborinates")
	for _, employee := range emp.GetSubordinates() {
		fmt.Println("id: ", employee.GetId(), "name: ", employee.GetName(), "managerID: ", employee.GetManagerID())

	}
}

func (system *System) printAllPrefixSearchName(prefix string) {
	for _, emp := range system.listEmployee {
		if strings.Index(emp.GetName(), prefix) == 0 {
			fmt.Println("id: ", emp.GetId(), "name: ", emp.GetName(), "managerID: ", emp.GetManagerID())
		}
	}
}
