package main

import "fmt"

func main() {
	arr := []int{2, 3, 1, 1, 4}
	fmt.Println(canJumpSol2(arr))
	fmt.Println(canJumpSol3(arr))
	fmt.Println(canJumpSol1(arr))
}

func canJumpSol3(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	// take as base condition
	maxJump := 0
	//run loop till end
	for i := 0; i < n; i++ {
		// check if maxJump is less than i then we are not able to each at i'th index then return false
		if maxJump < i {
			return false
		}
		maxJump = max(maxJump, i+nums[i]) // calculate new maxJump at every points.
	}
	return true
}

func canJumpSol2(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}
	maxJump := 0
	// run loop till maxJum and calculate maxjump every time, if there is no way to reach end and jour jump ends then
	//can not reach at end.
	for i := 0; i <= maxJump; i++ {

		maxJump = max(maxJump, i+nums[i])
		// check if we reached at end of array means return true.
		if maxJump >= n-1 {
			return true
		}
	}
	//we traversed all jump but not reach at the end of array then return false.
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJumpSol1(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	dp := make([]bool, n)

	dp[0] = true

	for i := 0; i < n; i++ {

		for j := 0; j < i; j++ {

			if nums[j] >= i-j {
				dp[i] = true
			}

		}
	}
	result := true
	for i := 0; i < n; i++ {
		if !dp[i] {
			result = false
		}
	}
	return result
}
