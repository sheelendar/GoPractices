package main

import (
	"fmt"
	"math"
)

/*
/*
You are given a network of n nodes labeled from 0 to n-1, and an array times which is a list of travel times as directed edges times[i] = (u(i), v(i), w(i)), where u(i) is the source node, v(i) is the target node, and w(i) is the time it takes for a signal to travel from source to target.

We will send a signal from a given node k. Return the minimum time it takes for all the n nodes to receive the signal. If it is impossible for all the n nodes to receive the signal, return -1.

Examples:
Input: times = [[0,1,1]], n = 2, k = 0
Output: 1
Graph:
0 ----> 1

	1

Input: times = [[0,1,1]], n = 2, k = 1
Output: -1
Graph:
0 ----> 1

	1

Explanation:
Since you need to start from node 1, you can't ever reach node 0

Input: times = [[1,0,3],[1,2,0],[2,3,1]], n = 4, k = 1
Output: 3
Explanation:
0 <---- 1 ----> 2  ----> 3

	3       0        1

To go from node 1 to 0, it costs 3 time units. Then, from 1 to 2 it costs 0, and from 2
*/
func main() {
	n := 4
	k := 1
	times := [][]int32{
		{1, 0, 3},
		{1, 2, 0},
		{2, 3, 1},
	}
	count := len(times)

	fmt.Print(miniTimeTotravel(times, count, n, k))
	// your solution here
}
func miniTimeTotravel(times [][]int32, count, n, k int) int32 {
	if n == 0 || k == -1 {
		return 0
	}
	graph := make([][]int32, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int32, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				graph[i][j] = 0
			} else {
				graph[i][j] = 999999
			}

		}
	}
	for i := 0; i < count; i++ {
		graph[times[i][0]][times[i][1]] = times[i][2]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				graph[i][j] = int32(math.Min(float64(graph[i][j]), float64(graph[i][k]+graph[k][j])))
			}
		}
	}
	maxTime := int32(0)
	for i := 0; i < n; i++ {
		if maxTime < graph[k][i] {
			maxTime = graph[k][i]
		}
	}
	if maxTime >= 999999 {
		return -1
	}
	return maxTime
}
