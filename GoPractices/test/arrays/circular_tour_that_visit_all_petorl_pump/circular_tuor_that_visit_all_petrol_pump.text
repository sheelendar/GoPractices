package main

import "fmt"

/*
Given information about N petrol pumps (say arr[]) that are present in a circular path. The information consists of the
distance of the next petrol pump from the current one (in arr[i][1]) and the amount of petrol stored in that petrol pump (in arr[i][0]).
Consider a truck with infinite capacity that consumes 1 unit of petrol to travel 1 unit distance. The task is to find the index of the
first starting point such that the truck can visit all the petrol pumps and come back to that starting point.
Note: Return -1 if no such tour exists.
Examples:

	Input: arr[] = {{4, 6}, {6, 5}, {7, 3}, {4, 5}}.
	Output: 1
	Explanation: If started from 1st index then a circular tour can be covered
*/
type petrolPump struct {
	petrol   int
	distance int
}

func main() {
	list := []petrolPump{{4, 6}, {6, 5}, {7, 3}, {4, 5}}
	size := len(list)
	fmt.Print("index from where tour complete tour ", findIndexToCompleteTour(list, size))
}

func findIndexToCompleteTour(list []petrolPump, size int) int {
	start := 0
	end := 1
	petrol := list[start].petrol - list[start].distance

	for start != end || petrol < 0 {

		for petrol < 0 && end != start {

			petrol = petrol - (list[start].petrol - list[start].distance)
			start = (start + 1) % size
			if start == 0 {
				return -1
			}
		}

		petrol = petrol + (list[end].petrol - list[end].distance)
		end = (end + 1) % size
	}
	return start
}
