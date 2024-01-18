/*You are given two integer arrays nums1 and nums2 of the same length n .
 Let's consider all cyclic t-shifts of nums1 as moving t elements from the end to the front of the array as stated below:

• For t = 0 , array nums1 = [nums1 [0], nums1 [1], nums1 [2], ..., numsl[n - 1]1

• For t = 1, array numsi = [nums1 [n - 1], nums1 [0]], nums1 [1], nums1 [2], nums1[3],
nums1 [n - 21

• For t = 2 , array nums1 = [nums1 [n - 2], nums1 [n - 1], nums1[0]l, nums1[1], nums1 [2],

nums1 [3], ..., nums1 [n - 3]
• For t = n-1, array nums1 = [nums1 [1], nums1 [2], nums1 [3], ..., nums1 [n - 1],
nums1 [01]

For example, the video below presents all cyclic t-shifts of the array nums = [5, 4, 3, 2, 1] :4 6 8 9 15 16

Note: If you are not able to see the video, use this link to access it.
Your task is the following: for each cyclic t-shift of nums1, calculate the sum of absolute differences with
nums2 . If the cyclic t-shift of nums1 is a new array numsShifted, the sum equals to
InumsShifted[0] - nums2[0]| + InumsShifted[1] - nums2[1]| + ... + InumsShifted[n - 1] - I
Return a list of all the sums sorted in non-descending order as the answer.

For nums1 = [1, 4, 2, 11] and nums2 = [10, 1, 8, 4] , the output should be solution(nums1,nums2) = [7, 13, 25, 25] .

Explanation:
	Let's consider all possible cyclic t-shifts of nums1 :

	• The 0 -shift is [1, 4, 2, 111, and the sum of absolute differences with nums2 = [10, 1, 8,
	41 :
	= 11 - 101 + 14 - 11 + 12 - 81 + 111 - 41 = 9 + 3 + 6+7 = 25.

	• The 1 -shift is [11, 1, 4, 21, and the sum of absolute differences with nums2 = [10, 1, 8,
	41 :
	= 111 - 101 + 11 - 11 + 14 - 81 + 12 - 41 = 1 + 0+4 + 2=7.

	• The 2 -shift is [2, 11, 1, 4], and the sum of absolute differences with nums2 = [10, 1, 8, 41:
	= 12 - 101 + 111 - 11 + 11 - 81 + 14 - 41 = 8 + 10 + 7 + 0 = 25

	• The 3 -shift is [4, 2, 11, 11, and the sum of absolute differences with nums2 = [10, 1, 8, 41 :
	= 14 - 101 + 12 - 11 + 111 - 81 + 11 - 41 = 6 + 1 + 3 + 3 = 13
	25
	The non-descending list of all the sums is [7, 13, 25, 25], which is the answer.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 4, 2, 11}
	nums2 := []int{10, 1, 8, 4}
	arr := solution(nums1, nums2)
	fmt.Println(arr)
}

func solution(nums1, nums2 []int) []int {
	n := len(nums1)
	var result []int
	arr := nums1
	for i := 0; i < n; i++ {
		if i != 0 {
			arr = rotateArray(arr)
		}
		sum := 0
		for j := 0; j < n; j++ {
			sum += ABS(arr[j] - nums2[j])
		}
		result = append(result, sum)
	}
	sort.Ints(result)
	return result
}
func ABS(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func rotateArray(arr []int) []int {
	n := len(arr)
	a := make([]int, n)

	for i := 0; i < n-1; i++ {
		a[i+1] = arr[i]
	}
	a[0] = arr[n-1]
	return a
}
