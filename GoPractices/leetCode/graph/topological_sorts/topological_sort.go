package main

import "fmt"

func main() {
	edges := make([][]int, 6)
	edges[0] = []int{5, 2}
	edges[1] = []int{5, 0}
	edges[2] = []int{4, 0}
	edges[3] = []int{4, 1}
	edges[4] = []int{2, 3}
	edges[5] = []int{3, 1}
	displayTopologicalOrder(edges)
}

func displayTopologicalOrder(edges [][]int) {
	size := len(edges)
	graph := make([][]int, size)

	for i := 0; i < size; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
	}
	displayGraph(graph)

	var stack []int
	visited := make([]bool, size)
	for i := 0; i < size; i++ {
		DFS(i, graph, &stack, visited)
	}

	fmt.Println()
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Print(stack[i])
		fmt.Print(" ")
	}
}

func DFS(n int, adj [][]int, stack *[]int, visited []bool) {

	if visited[n] {
		return
	}
	for j := 0; j < len(adj[n]); j++ {
		if !visited[adj[n][j]] {
			DFS(adj[n][j], adj, stack, visited)
		}
	}
	visited[n] = true
	*stack = append(*stack, n)
}

func displayGraph(graph [][]int) {
	for i := 0; i < len(graph); i++ {
		fmt.Print(i)
		fmt.Print(": ")
		for j := 0; j < len(graph[i]); j++ {
			fmt.Print(graph[i][j])
			fmt.Print(",")
		}
		fmt.Println()
	}
}
