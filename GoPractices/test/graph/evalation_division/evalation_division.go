package main

import "fmt"

/*
You are given an array of variable pairs equations and an array of real numbers values, where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i]. Each Ai or Bi is a string that represents a single variable.

You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query where you must find the answer for Cj / Dj = ?.

Return the answers to all queries. If a single answer cannot be determined, return -1.0.

Note: The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.

Note: The variables that do not occur in the list of equations are undefined, so the answer cannot be determined for them.

Example 1:

Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
Explanation:
Given: a / b = 2.0, b / c = 3.0
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
note: x is undefined => -1.0

https://leetcode.com/problems/evaluate-division/description/?envType=study-plan-v2&envId=top-interview-150
*/

func main() {
	//equations := [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	//values := []float64{1.5, 2.5, 5.0}
	//queries := [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}

	equations := [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}
	values := []float64{3.0, 4.0, 5.0, 6.0}
	queries := [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}
	result := calcEquation(equations, values, queries)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i], "  ")
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	m := len(equations)
	for i := 0; i < m; i++ {
		a := equations[i][0] //  fetch both vertics
		b := equations[i][1]
		adj, ok := graph[a] // put vertic a->b
		if ok {
			adj[b] = values[i]
		} else {
			graph[a] = map[string]float64{b: values[i]}
		}
		adj, ok = graph[b] // put vertic b->a
		if ok {
			adj[a] = 1 / values[i]
		} else {
			graph[b] = map[string]float64{a: 1 / values[i]}
		}
	}
	n := len(queries)
	result := make([]float64, n)
	for i := 0; i < n; i++ {
		a := queries[i][0]
		b := queries[i][1]
		result[i] = findValue(graph, a, b, make(map[string]bool)) // fetch result for each query
	}
	return result
}

func findValue(graph map[string]map[string]float64, a, b string, visited map[string]bool) float64 {
	if _, ok := graph[a]; !ok { // if a is not found then return -1.0
		return -1.0
	}
	if a == b { // if no both are equl return 1.0
		return 1.0
	}
	visited[a] = true
	adj := graph[a]
	if v, ok := adj[b]; ok { // if b is not found then return -1.0
		return v
	}
	for k, v := range adj { //  check all adjecet of the vertics and return if found value
		if !visited[k] {
			product := findValue(graph, k, b, visited)
			if product != -1.0 {
				return v * product
			}
		}
	}
	return -1.0
}
