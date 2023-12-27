package main

import "fmt"

func main() {
	arr := []int{8, 9, 8, 1, 3, 4, 6, 5, 7, 2}
	target := 9
	fmt.Println(findPairSumIndex(arr, target))

}

func findPairSumIndex(arr []int, target int) [][]int {
	size := len(arr)
	if size < 2 {
		return nil
	}
	dp := make(map[int][]int)
	var res [][]int
	for i := 0; i < size; i++ {
		if arr[i] > target {
			continue
		}
		if indexes, ok := dp[arr[i]]; ok {

			for j := 0; j < len(indexes); j++ {
				res = append(res, []int{indexes[j], i})
			}

		} else {

			num := target - arr[i]
			dp[num] = append(dp[num], i)

		}
	}
	return res
}
