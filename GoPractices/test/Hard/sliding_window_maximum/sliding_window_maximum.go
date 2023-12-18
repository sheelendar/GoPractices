/*
You are given an array of integers nums, there is a sliding window of size k which is moving from
the very left of the array to the very right. You can only see the k numbers in the window.
Each time the sliding window moves right by one position.

Return the max sliding window.

Example 1:
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6 7       3
1 [3  -1  -3] 5  3  6  7       3
1  3 [-1  -3  5] 3  6  7       5
1  3  -1 [-3  5  3] 6  7       5
1  3  -1  -3 [5  3  6] 7       6
1  3  -1  -3  5 [3  6  7]      7

Example 2:
Input: nums = [1], k = 1
Output: [1]
*/
package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := maxSlidingWindow(nums, k)
	fmt.Println(result)
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var queue []int
	n := len(nums)
	for i := 0; i < n; i++ {
		size := len(queue) // delete if new value is bigger then previous values then biggger will be ans
		for size > 0 && nums[queue[size-1]] <= nums[i] {
			queue = queue[:size-1]
			size = size - 1
		}
		queue = append(queue, i) //add new value add position

		if i >= (k - 1) { // check if i is equal or grater than k and add in answers
			result = append(result, nums[queue[0]])
		}

		if queue[0] <= i-(k-1) { // check if bigger value index is out of k size window then delete that
			queue = queue[1:]
		}
	}
	return result
}
