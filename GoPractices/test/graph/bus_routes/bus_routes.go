/*
You are given an array routes representing bus routes where routes[i] is a bus route that the ith bus repeats forever.

	For example, if routes[0] = [1, 5, 7], this means that the 0th bus travels in the sequence 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... forever.

You will start at the bus stop source (You are not on any bus initially), and you want to go to the bus stop target. You can travel between bus stops by buses only.

Return the least number of buses you must take to travel from source to target. Return -1 if it is not possible.

Example 1:

Input: routes = [[1,2,7],[3,6,7]], source = 1, target = 6
Output: 2
Explanation: The best strategy is take the first bus to the bus stop 7, then take the second bus to the bus stop 6.

Example 2:

Input: routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
Output: -1
*/
package main

import "fmt"

func main() {
	routes := [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}
	source := 15
	target := 12
	// output= -1
	fmt.Println(numBusesToDestination(routes, source, target))
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	graph := make(map[int][]int)
	if source == target {
		return 0
	}
	n := len(routes)
	for i := 0; i < n; i++ {
		for j := 0; j < len(routes[i]); j++ {
			graph[routes[i][j]] = append(graph[routes[i][j]], i)
		}
	}

	for k, v := range graph {
		fmt.Print(k, " - ")
		fmt.Println(v)
	}

	var queue []int
	queue = append(queue, source)
	count := 0
	visited := make(map[int]bool)

	for len(queue) > 0 {
		size := len(queue)
		count++
		for i := 0; i < size; i++ {
			stop := queue[0]
			queue = queue[1:]

			buses := graph[stop]
			for _, bus := range buses {
				if !visited[bus] {
					for j := 0; j < len(routes[bus]); j++ {
						if routes[bus][j] == target {
							return count
						}
						queue = append(queue, routes[bus][j])
						visited[bus] = true
					}
				}

			}
		}
	}
	return -1
}
