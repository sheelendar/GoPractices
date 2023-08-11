package main

import "fmt"

func main() {
	m := [][]int{
		{1, 2, 3},
		{5, 6, 7},
		{8, 9, 10},
		{11, 12, 13},
	}
	result := SpiralCopy(m)
	result1 := SpiralCopy1(m)
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		fmt.Print(" ")
	}
	println()
	for i := 0; i < len(result1); i++ {
		fmt.Print(result1[i])
		fmt.Print(" ")
	}
}

func SpiralCopy(m [][]int) []int {
	ROW := len(m)
	COL := len(m[0])
	r := 0
	c := 0
	var result []int
	for r < ROW && c < COL {
		for i := c; i < COL; i++ {
			result = append(result, m[r][i])
		}
		r++
		for i := r; i < ROW; i++ {
			result = append(result, m[i][COL-1])
		}
		COL--
		for i := COL - 1; i >= c; i-- {
			result = append(result, m[ROW-1][i])
		}
		ROW--
		for i := ROW - 1; i >= r; i-- {
			result = append(result, m[i][c])
		}
		c++
	}
	return result
}

func SpiralCopy1(m [][]int) []int {
	ROW := len(m)
	COL := len(m[0])
	var result []int
	sprialHelper(m, 0, 0, ROW, COL, &result)
	return result
}

func sprialHelper(m [][]int, r, c, ROW, COL int, result *[]int) {
	if r < 0 || r >= ROW || c < 0 || c >= COL || m[r][c] == 999999 {
		return
	}
	*result = append(*result, m[r][c])
	m[r][c] = 999999
	if r == 0 || m[r-1][c] == 999999 {
		sprialHelper(m, r, c+1, ROW, COL, result)
	}
	sprialHelper(m, r+1, c, ROW, COL, result)
	sprialHelper(m, r, c-1, ROW, COL, result)
	sprialHelper(m, r-1, c, ROW, COL, result)
}
