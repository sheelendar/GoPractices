package main

import (
	"fmt"
)

type Node struct {
	r    int
	c    int
	cost int
}

func main() {
	sr := 0
	sc := 0
	tr := 2
	tc := 0
	/*grid := [][]int{{1, 1, 1, 1},
	{0, 0, 0, 1},
	{1, 0, 1, 1}}*/
	grid := [][]int{{1, 1, 1, 1},
		{0, 0, 0, 1},
		{1, 1, 1, 1}}
	fmt.Println(sortestPathBetWeenTwoNode(grid, sr, sc, tr, tc))
}
func sortestPathBetWeenTwoNode(grid [][]int, sr, sc, tr, tc int) int {
	R := len(grid)
	C := len(grid[0])
	node := Node{r: sr, c: sc, cost: 0}
	queue := []Node{node}
	result := make([][]bool, R)
	for i := 0; i < R; i++ {
		result[i] = make([]bool, C)
	}
	for len(queue) > 0 {
		n := queue[0]
		result[n.r][n.c] = true
		queue = queue[1:]
		if n.r == tr && n.c == tc {
			return n.cost
		}
		if check(n.r-1, n.c, R, C) && grid[n.r-1][n.c] == 1 && !result[n.r-1][n.c] {
			node := Node{r: n.r - 1, c: n.c, cost: n.cost + 1}
			queue = append(queue, node)
		}
		if check(n.r, n.c-1, R, C) && grid[n.r][n.c-1] == 1 && !result[n.r][n.c-1] {
			node := Node{r: n.r, c: n.c - 1, cost: n.cost + 1}
			queue = append(queue, node)
		}
		if check(n.r+1, n.c, R, C) && grid[n.r+1][n.c] == 1 && !result[n.r+1][n.c] {
			node := Node{r: n.r + 1, c: n.c, cost: n.cost + 1}
			queue = append(queue, node)
		}
		if check(n.r, n.c+1, R, C) && grid[n.r][n.c+1] == 1 && !result[n.r][n.c+1] {
			node := Node{r: n.r, c: n.c + 1, cost: n.cost + 1}
			queue = append(queue, node)
		}
	}
	return -1
}
func check(sr, sc, R, C int) bool {
	if sr >= 0 && sr < R && sc >= 0 && sc < C {
		return true
	}
	return false

}
