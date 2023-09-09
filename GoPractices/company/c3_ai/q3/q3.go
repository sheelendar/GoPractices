package main

import (
	"fmt"
)

func main() {
	//arr := []int32{2, 2, 2, 2}
	// arr := []int32{3, 3, 3, 3, 3, 1, 3}
	arr := []int32{1, 1, 1}
	socialGraphs(arr)
}

func socialGraphs(counts []int32) {
	//persons := len(counts)

	groupMap := make(map[int32][]int32)

	for i, grp := range counts {
		if _, exists := groupMap[grp]; exists {
			groupMap[grp] = append(groupMap[grp], int32(i))
		} else {
			groupMap[grp] = []int32{int32(i)}
		}
	}

	var sortedGroups []int32
	for grp := range groupMap {
		sortedGroups = append(sortedGroups, int32(grp))
	}
	for _, key := range sortedGroups {
		indices := groupMap[key]
		size := len(indices) / int(key)

		for i := 0; i < size; i++ {
			for j := 0; j < int(key); j++ {
				if len(indices) > 0 {
					x := indices[0]
					indices = indices[1:]
					fmt.Print(x, " ")
				}
			}
			fmt.Println()
		}
	}
}
