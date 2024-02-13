package main

import "fmt"

/*

    A NxM Maze matrix and another integer K is given ,Matrix contains only 0's and 1's. 0=empty space while 1= obstacle .
	One has to move from (0,0) to (N-1,M-1) with following constraints-
    One can move from (i,j) to (i+k,j) if there is no obstacle between (i,j) and (i+k,j) ,
    (i,j) to (i-k,j) if there is no obstacle between (i,j) and (i-k,j) ,
    (i,j) to (i,j+k) if there is no obstacle between (i,j) and (i,j+k) ,
    (i,j) to (i,j-k) if there is no obstacle between (i,j) and (i,j-k) .
    Return minimum moves for this .
    If one cannot reach destination , return -1 .
    Constraints - (1<=N,M,k<=100)

Eg - For 2x3 [[0,0,0],[1,0,0]] -> Output 2
For 3x3 [[0,1,0],[0,1,0],[0,0,0]] -> Output 2

*/

func main() {
	/*maza := [][]int{
	{0, 1, 0},
	{0, 1, 0},
	{1, 0, 0}} */

	maza := [][]int{
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0}}

	/*maza := [][]int{
	{0, 0, 0},
	{1, 0, 0}}*/
	k := 5
	fmt.Println(findMininumSteps(maza, k))
}

func findMininumSteps(maza [][]int, k int) int {
	m := len(maza)
	n := len(maza[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var queue []Pair
	queue = append(queue, Pair{row: 0, col: 0})

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			pair := queue[0]
			queue = queue[1:]
			visited[pair.row][pair.col] = true
			if pair.row >= m-1 && pair.col >= n-1 {
				return pair.count
			}
			for j := 1; j <= k; j++ {
				if pair.row+j < m && maza[pair.row+j][pair.col] == 0 && !visited[pair.row+j][pair.col] {
					queue = append(queue, Pair{row: pair.row + j, col: pair.col, count: pair.count + 1})
				} else {
					break
				}
			}
			for j := 1; j <= k; j++ {
				if pair.row-j >= 0 && maza[pair.row-j][pair.col] == 0 && !visited[pair.row-j][pair.col] {
					queue = append(queue, Pair{row: pair.row - j, col: pair.col, count: pair.count + 1})
				} else {
					break
				}
			}

			for j := 1; j <= k; j++ {
				if pair.col+j < n && maza[pair.row][pair.col+j] == 0 && !visited[pair.row][pair.col+j] {
					queue = append(queue, Pair{row: pair.row, col: pair.col + j, count: pair.count + 1})
				} else {
					break
				}
			}

			for j := 1; j <= k; j++ {
				if pair.col-j >= 0 && maza[pair.row][pair.col-j] == 0 && !visited[pair.row][pair.col-j] {
					queue = append(queue, Pair{row: pair.row, col: pair.col - j, count: pair.count + 1})
				} else {
					break
				}
			}
		}
	}

	return -1
}

type Pair struct {
	row   int
	col   int
	count int
}
