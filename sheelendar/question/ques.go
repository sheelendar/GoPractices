package main

import (
	"fmt"
	"sort"
)

/*Tesco has around 3200 stores and more than 10% of the stores have around 800 colleagues each.

In a store, a colleague can work for multiple departments. Here are shifts of a colleague in various departments:

In Bakery department: From 8 to 10
In Checkout department: From 10 to 12  (13-14)
In Diary department: From 14 to 19
Given the above split of shifts, provide an API/method to return the different shifts of a colleague for the day after merging contiguous shifts.
 This will be exposed to the colleague in different UI and help them plan their day accordingly.

His shift timings in this case are 8 to 12 and 14 to 19.
*/
//9 - 11 ,15-20

//7,

func main() {
	departmentShifts := []DepartmentShift{DepartmentShift{8, 10, "bakery"}, DepartmentShift{10, 12, "checkput"}, DepartmentShift{14, 19, "deary"}}
	employeetime := [][]int{{7, 11}, {13, 15}, {17, 20}}
	//employeetime := [][]int{{8, 12}, {14, 19}}
	// he can work in Bakery 8-10 and checkoout : 10- 12,Diary 14,19

	result := findTiming(departmentShifts, employeetime)
	for _, res := range result {
		fmt.Println("Can work in ", res)
	}

}

func findTiming(departmentShifts []DepartmentShift, empTime [][]int) []string {

	n := len(departmentShifts)
	m := len(empTime)
	if n == 0 || m == 0 {
		return []string{}
	}
	sort.Slice(departmentShifts, func(i, j int) bool {
		return departmentShifts[i].start < departmentShifts[j].start
	})
	sort.Slice(empTime, func(i, j int) bool {
		return empTime[i][0] < empTime[j][0]
	})
	var result []string
	departmentHash := make(map[string]bool)

	for i := 0; i < m; i++ {
		empShift := empTime[i]

		for _, department := range departmentShifts {

			if empShift[0] <= department.start && !departmentHash[department.name] {
				result = append(result, fmt.Sprint("department :", department.name, " - ", empShift[0], "- ", Min(department.end, empShift[1])))
				departmentHash[department.name] = true
				empShift[0] = department.end
			}
			if empShift[0] >= empShift[1] {
				break
			}
		}

	}
	return result
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type DepartmentShift struct {
	start int
	end   int
	name  string
}
