package main

import "fmt"

func main() {
	graph := make(map[int][]int)
	//edges := [][]int{{1, 0}, {0, 2}, {0, 3}, {3, 4}}
	edges := [][]int{{1, 0}, {0, 2}, {0, 3}, {3, 4}, {5, 6}}
	numOfNodes := make(map[int]bool)
	// put into a graph as key will be node and values will be connected nodes.
	for i := 0; i < len(edges); i++ {
		// this is undirected graph so we put both edges into graph v->u and u->v.
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])

		// check unique number of nodes into graph
		if !numOfNodes[edges[i][0]] {
			numOfNodes[edges[i][0]] = true
		}
		if !numOfNodes[edges[i][1]] {
			numOfNodes[edges[i][1]] = true
		}
	}
	//fmt.Println(isGraphTree(graph, len(edges), len(numOfNodes)))
	fmt.Print(findConnectedComponents(graph, len(numOfNodes)))
}

func findConnectedComponents(graph map[int][]int, numOfNodes int) int {
	count := 0
	isVisited := make([]bool, numOfNodes)

	for n, _ := range graph {
		if !isVisited[n] {
			DFS(n, graph, isVisited)
			count++
		}
	}
	return count
}

func DFS(n int, graph map[int][]int, visited []bool) {
	if visited[n] {
		return
	}
	visited[n] = true
	adj := graph[n]

	for i := 0; i < len(adj); i++ {
		DFS(adj[i], graph, visited)
	}
}
