package main

/*
The Floyd Warshall Algorithm is for solving all pairs of shortest-path problems.
The problem is to find the shortest distances between every pair of vertices in a given edge-weighted directed Graph.
It is an algorithm for finding the shortest path between all the pairs of vertices in a weighted graph.
This algorithm follows the dynamic programming approach to find the shortest path.
Input:  graph[][] = { {0,   5,  INF, 10},

	{INF,  0,  3,  INF},
	{INF, INF, 0,   1},
	{INF, INF, INF, 0} }

which represents the following graph

	       10
	(0)——----------->(3)
	  |              /|\
	5 |               |  1
	  |               |
	 \|/             |
	(1)——---------->(2)
	        3

Output: Shortest distance matrix

	   0        5      8       9
	INF       0      3       4
	INF     INF    0       1
	INF     INF    INF    0
*/
const INF = 999999

func main() {
	graph := [][]int64{{0, 5, INF, 10},
		{INF, 0, 3, INF},
		{INF, INF, 0, 1},
		{INF, INF, INF, 0}}
	floydWarshallAllShortestPathAlgo(graph, len(graph))
}

func floydWarshallAllShortestPathAlgo(graph [][]int64, size int) {
	for k := 0; k < size; k++ {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if graph[i][j] > graph[i][k]+graph[k][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
				}
			}
		}
	}
	displayGraph(graph, size)
}

func displayGraph(graph [][]int64, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if graph[i][j] == INF {
				print("a")
			} else {
				print(graph[i][j])
			}
			print(" ")
		}
		println()
	}
}
