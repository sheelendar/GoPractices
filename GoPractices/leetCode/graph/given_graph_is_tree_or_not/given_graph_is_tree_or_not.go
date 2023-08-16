package main

import "fmt"

// https://www.geeksforgeeks.org/check-given-graph-tree/
/*

Check if a given graph is tree or not
Write a function that returns true if a given undirected graph is a tree and false otherwise. For example, the following graph is a tree.

An undirected graph is a tree if it has the following properties.
    There is no cycle.
    The graph is connected.

However if we observe carefully the definition of tree and its structure we will deduce that
	if a graph is connected and has n – 1 edges exactly then the graph is a tree.
	Proof:
	Since we have assumed our graph of n nodes to be connected, it must have at least n – 1 edges inside it.
	Now if we try to add one more edge than the n – 1 edges already the graph will end up forming a cycle and
	thus will not satisfy the definition of tree. Therefore, it is necessary for a connected graph to have exactly n – 1
	edges to avoid forming cycle.
*/

func main() {
	graph := make(map[int][]int)
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
	fmt.Println(isGraphTree(graph, len(edges), len(numOfNodes)))
	//displayGraphEdge(graph)
}

func isGraphTree(graph map[int][]int, edges, numOfNodes int) bool {
	// if number of edges is more than number of nodes then it is not a tree. then cycle will be there.
	//so it is property to make it a tree edges should be less than nodes.
	if edges >= numOfNodes {
		return false
	}
	isVisited := make([]bool, numOfNodes)
	checkConnectedNodes(0, graph, isVisited)
	for i := 0; i < len(isVisited); i++ {
		// if we found disconnected node then it is not a tree
		if !isVisited[i] {
			return false
		}
	}
	return true
}

func checkConnectedNodes(n int, graph map[int][]int, isVisited []bool) {
	if isVisited[n] {
		return
	}
	isVisited[n] = true
	adj, _ := graph[n]
	for i := 0; i < len(adj); i++ {
		checkConnectedNodes(adj[i], graph, isVisited)
	}
}

func displayGraphEdge(graph map[int][]int) {
	for k, v := range graph {
		fmt.Print(k)
		fmt.Print(" : ")
		for i := 0; i < len(v); i++ {
			fmt.Print(v[i])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
