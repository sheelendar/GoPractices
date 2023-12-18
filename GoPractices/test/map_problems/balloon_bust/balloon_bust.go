package main

import "fmt"

/*
Given heights at which balloons are placed.
You get to shoot an arrow that goes from (L -> R), starting at any height H.
If height of arrow = height of balloon, then the balloon pops and the height of arrow drops by 1.
Minimum number of arrows required to pop all balloons?

[9,8, 6, 5, 7] => 1,1,2,

[2,3] => 2
*/

func main() {
	arr := []int{9, 8, 6, 5, 7}
	fmt.Println(burstingBalloons(arr, len(arr)))
}

func burstingBalloons(arr []int, n int) int {
	arrowsAtHeight := make(map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := arrowsAtHeight[arr[i]]; ok {
			// Destroying Balloon, so frequency will decrease.
			arrowsAtHeight[arr[i]]--

			if arrowsAtHeight[arr[i]] == 0 {
				delete(arrowsAtHeight, arr[i])
			}
		}
		// Frequency of Arrow at 1 less height will increase.
		arrowsAtHeight[arr[i]-1]++
	}
	ans := 0
	for _, freq := range arrowsAtHeight {
		ans += freq
	}
	return ans
}
