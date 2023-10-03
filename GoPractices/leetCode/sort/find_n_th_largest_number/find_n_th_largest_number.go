package main

import "fmt"

func main() {
	arr := []int{3, 1, 2, 4}
	k := 2
	fmt.Println(findKthLargest(arr, k))

}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	size := n
	//res := 0
	for i := 0; i < k; i++ {
		nums = heapify(nums, size)
		//res = nums[0]
		size--
		nums[0], nums[size] = nums[size], nums[0]
	}
	return nums[size]
}
func heapify(arr []int, n int) []int {
	for parant := (n - 1) / 2; parant >= 0; parant-- {
		right := parant*2 + 1
		left := parant * 2
		if right < n && arr[parant] < arr[right] {
			arr[parant], arr[right] = arr[right], arr[parant]
		}
		if arr[parant] < arr[left] {
			arr[parant], arr[left] = arr[left], arr[parant]
		}
	}
	return arr
}
