package main

import (
	"fmt"
	"sort"
)

func main() {
	//envelopes := [][]int{{4, 3}, {5, 3}, {5, 6}, {1, 2}}
	//output:=3
	envelopes := [][]int{{3, 6}, {5, 4}, {4, 8}, {6, 9}, {10, 7}, {12, 12}}
	//output:=4
	fmt.Println(maxEnvelopes(envelopes))
}

type Envelop struct {
	height int
	width  int
}

func maxEnvelopes(arr [][]int) int {
	m := len(arr)
	if m == 0 {
		return 0
	}
	var envelops []Envelop
	for i := 0; i < m; i++ {
		envelops = append(envelops, Envelop{height: arr[i][0], width: arr[i][1]})
	}
	sort.Slice(envelops, func(i, j int) bool {
		return envelops[i].height*envelops[i].width > envelops[j].height*envelops[j].width
	})
	size := len(envelops)
	dp := make([]int, size)
	dp[0] = 1
	for i := 1; i < size; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelops[i].height < envelops[j].height && envelops[i].width < envelops[j].width {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[size-1]
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
