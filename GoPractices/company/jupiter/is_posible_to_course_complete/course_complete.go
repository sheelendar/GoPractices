package main

/*


////numCourses = 2, prerequisites = [[1,0],[0,1]] - false
// numCoursew = 2, prerequisites = [[1,0]] - true
// numCoirse=3, pre = [[1,0],[0,2]] - true

*/

func main() {
	// 1-> 0
	// 0 >1

}

/*isSamePath[0]=true
visited[0]=true

isSamePath[1]=true
visited[1]=true
*/

func findCycle(node int, graph map[int][]int, visited, isSamePath []bool) bool {
	if isSamePath[node] {
		return true
	}
	if visited[node] {
		return false
	}
	isSamePath[node] = true
	visited[node] = true

	adj := graph[node]
	for i := 0; i < len(adj); i++ {
		if !visited[adj[i]] {
			res := findCycle(adj[i], graph, visited, isSamePath)
			if res {
				return res
			}
		}
	}

	isSamePath[node] = false
	return false
}
