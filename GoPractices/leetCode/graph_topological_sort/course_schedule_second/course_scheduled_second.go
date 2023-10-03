package main

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates
that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return the ordering of courses you should take to finish all courses.
If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have
finished course 0. So the correct course order is [0,1].

Example 2:

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have
finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

*/

func main() {

}

func findOrder(numCourses int, prerequisites [][]int) []int {
	n := len(prerequisites)
	var result []int
	if n == 0 {
		for i := 0; i < numCourses; i++ {
			result = append(result, i)
		}
		return result
	}
	dp := make(map[int][]int)
	for i := 0; i < n; i++ {
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		dp[a] = append(dp[a], b)
	}
	visited := make([]bool, numCourses)
	samePath := make([]bool, numCourses)
	var stack []int
	for node, _ := range dp {
		isCycle := DFS(node, dp, visited, samePath, &stack)
		if isCycle {
			return result
		}
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			stack = append(stack, i)
		}
	}
	for i := 0; i < len(stack); i++ {
		result = append(result, stack[i])
	}
	return result
}

func DFS(node int, graph map[int][]int, visited, samePath []bool, stack *[]int) bool {
	if samePath[node] {
		return true
	}
	if visited[node] {
		return false
	}
	visited[node] = true
	samePath[node] = true

	adj := graph[node]
	for i := 0; i < len(adj); i++ {
		if DFS(adj[i], graph, visited, samePath, stack) {
			return true
		}
	}
	*stack = append(*stack, node)
	samePath[node] = false
	return false
}
