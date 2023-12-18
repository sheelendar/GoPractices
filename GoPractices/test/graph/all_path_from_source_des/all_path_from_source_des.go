/*
Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths from node 0 to node n - 1 and return them in any order.
The graph is given as follows: graph[i] is a list of all nodes you can visit from node i (i.e., there is a directed edge from node i to node graph[i][j]).

Input: graph = [[1,2],[3],[3],[]]
Output: [[0,1,3],[0,2,3]]
Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
*/
package main

import "fmt"

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}

type Path struct {
	node int
	arr  []int
}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	var result [][]int
	var queue []Path
	arr := []int{0}

	queue = append(queue, Path{node: 0, arr: arr})
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		fmt.Println(len(queue))
		if item.node == n-1 {
			fmt.Println(item.arr)
			fmt.Print(item.node)
			result = append(result, item.arr)
			continue
		}
		for j := 0; j < len(graph[item.node]); j++ {
			a := make([]int, len(item.arr))
			copy(a, item.arr)
			a = append(a, graph[item.node][j])
			newItem := Path{node: graph[item.node][j], arr: a}
			queue = append(queue, newItem)
			fmt.Println(len(queue))
		}
	}
	return result
}
