package main

/*
Given an array arr[] where each element represents the max number of steps that can be made forward from that index.
The task is to find the minimum number of jumps to reach the end of the array starting from index 0. If the end isn’t
reachable, return -1.
Examples:

	Input: arr[] = {1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9}
	Output: 3 (1-> 3 -> 9 -> 9)
	Explanation: Jump from 1st element to 2nd element as there is only 1 step.
	Now there are three options 5, 8 or 9. I
	f 8 or 9 is chosen then the end node 9 can be reached. So 3 jumps are made.

	Input:  arr[] = {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	Output: 10
	Explanation: In every step a jump is needed so the count of jumps is 10.
*/

func main() {
	arr := []int64{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9}
	//arr := []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	println("minimum of jumps", minNumberOfJumps(arr, len(arr)))
}

func minNumberOfJumps(arr []int64, size int) int64 {
	if size == 0 {
		return 0
	}
	if arr[0] == 0 {
		return -1
	}
	jumps := make([]int64, size+1)
	for i := 1; i < size; i++ {
		jumps[i] = int64(size + 1)
	}
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if arr[j]+int64(j) >= int64(i) && jumps[i] > jumps[j]+1 {
				jumps[i] = jumps[j] + 1
			}
		}
	}
	return jumps[size-1]
}
