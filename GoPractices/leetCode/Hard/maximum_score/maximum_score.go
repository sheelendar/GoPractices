package main

import "fmt"

/*
You are given an array of integers nums (0-indexed) and an integer k.
The score of a subarray (i, j) is defined as min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1). A good subarray is a subarray where i <= k <= j.
Return the maximum possible score of a good subarray.

Example 1:
Input: nums = [1,4,3,7,4,5], k = 3
Output: 15
Explanation: The optimal subarray is (1, 5) with a score of min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15.

Example 2:
Input: nums = [5,5,4,5,4,1,1,1], k = 0
Output: 20
Explanation: The optimal subarray is (0, 4) with a score of min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20.
*/
func main() {
	nums := []int{1, 4, 3, 7, 4, 5}
	k := 3
	fmt.Println(maximumScore(nums, k))
}

func maximumScore(nums []int, k int) int {
	max_area := 0
	var stack []int
	size := len(nums)
	i := 0
	for i < size {
		n := len(stack)
		if n == 0 || nums[stack[n-1]] <= nums[i] {
			stack = append(stack, i)
			i++
		} else {
			m := stack[n-1]
			stack = stack[:n-1]
			if len(stack) == 0 && i >= k && 0 <= k {
				max_area = Max(max_area, nums[m]*i)
			} else if stack[len(stack)-1]-1 <= k && i >= k {
				max_area = Max(max_area, nums[m]*(i-stack[len(stack)-1]-1))
			}
		}
	}

	for len(stack) > 0 {
		n := len(stack)
		m := stack[n-1]
		stack = stack[:n-1]
		if len(stack) == 0 && i >= k && 0 <= k {
			max_area = Max(max_area, nums[m]*i)
		} else if stack[len(stack)-1] <= k && i >= k {
			max_area = Max(max_area, nums[m]*(i-stack[len(stack)-1]-1))
		}
	}
	return max_area
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaximumScore(A []int, k int) int {
	res := A[k]
	mini := A[k]
	i := k
	j := k
	n := len(A)
	for i > 0 || j < n-1 {
		if i == 0 {
			j++
		} else if j == n-1 {
			i--
		} else if A[i-1] < A[j+1] {
			j++
		} else {
			i++
		}
		mini = Min(mini, Min(A[i], A[j]))
		res = Max(res, mini*(j-i+1))
	}
	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
