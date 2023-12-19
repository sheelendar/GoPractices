package main

import (
	"fmt"
	"sort"
)

type Box struct {
	width  int
	length int
	height int
}

func main() {
	length := []int{4, 1, 4, 10}
	width := []int{6, 2, 5, 12}
	height := []int{7, 3, 6, 32}
	fmt.Print(maxHeight(length, width, height))
}

func maxHeight(length []int, width []int, height []int) int {
	n := len(length)
	if n == 0 {
		return 0
	}
	var boxes []Box
	for i := 0; i < n; i++ {
		boxes = append(boxes, Box{length: width[i], width: height[i], height: length[i]})
		boxes = append(boxes, Box{width: width[i], height: height[i], length: length[i]})
		boxes = append(boxes, Box{height: width[i], length: height[i], width: length[i]})
	}
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].length*boxes[i].width > boxes[j].length*boxes[j].width
	})
	size := len(boxes)
	dp := make([]int, size)
	maxH := 0
	//dp[0] = boxes[0].height
	for i := 0; i < size; i++ {
		dp[i] = boxes[i].height
		for j := 0; j < i; j++ {
			if boxes[i].length < boxes[j].length && boxes[i].width < boxes[j].width {
				dp[i] = Max(dp[i], dp[j]+boxes[i].height)
			}
		}
		maxH = Max(maxH, dp[i])
	}
	return maxH
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
