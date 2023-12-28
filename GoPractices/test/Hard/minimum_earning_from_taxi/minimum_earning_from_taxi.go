package main

import (
	"fmt"
	"sort"
)

func main() {
	//n := 20
	//rides := [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}
	n := 10
	rides := [][]int{{2, 3, 6}, {8, 9, 8}, {5, 9, 7}, {8, 9, 1}, {2, 9, 2}, {9, 10, 6}, {7, 10, 10}, {6, 7, 9}, {4, 9, 7}, {2, 3, 1}}
	fmt.Println(maxTaxiEarnings(n, rides))

	/*
		2  -  {2 1}
		4  -  {12 5}
		5  -  {11 4}
		6  -  {10 1}
		7  -  {13 3}
		8  -  {2 1}
		9  -  {7 1}

		0 2 2 14 14 14 27 27 34
	*/
}
func maxTaxiEarnings(n int, rides [][]int) int64 {
	dp := make(map[int]Pair)
	price := make([]int64, n+1)

	sort.Slice(rides, func(i, j int) bool {
		return rides[i][0] < rides[j][0]
	})

	for i := 0; i < len(rides); i++ {
		cost := rides[i][1] - rides[i][0] + rides[i][2]
		dp[rides[i][0]] = Pair{cost: int64(cost), window: int64(rides[i][1] - rides[i][0])}
	}
	window := int64(0)
	maxPrice := int64(0)
	for i := 1; i < n; i++ {

		if pair, ok := dp[i]; ok {

			c := int64(0)
			if i-int(window) >= 0 {
				c = price[i-int(window)]
			}
			if c+pair.cost > price[i-1] {
				price[i] = c + pair.cost
				window = Max(window, pair.window)
			} else {
				price[i] = price[i-1]
			}

		} else {
			price[i] = price[i-1]
		}
		if window > 0 {
			window--
		}
		maxPrice = Max(maxPrice, price[i])
	}
	for i := 1; i < n; i++ {
		fmt.Print(price[i], " ")
	}
	return maxPrice
}
func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type Pair struct {
	cost   int64
	window int64
}
