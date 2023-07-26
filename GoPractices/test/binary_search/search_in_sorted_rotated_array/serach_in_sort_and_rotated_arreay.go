package main

import "fmt"

func main() {
	arr := []int64{4, 5, 6, 7, 8, 9, 1, 2, 3}
	k := int64(3)
	index := searchInRotatedArray(arr, k, 0, len(arr)-1)
	fmt.Println("find the value", index)
}

func searchInRotatedArray(arr []int64, k int64, start, end int) interface{} {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if k == arr[mid] {
		return mid
	}

	if arr[start] <= arr[mid] {
		if k >= arr[start] && k <= arr[mid] {
			return searchInRotatedArray(arr, k, start, mid-1)
		}
		return searchInRotatedArray(arr, k, mid+1, end)

	}
	if k >= arr[mid] && k <= arr[end] {
		return searchInRotatedArray(arr, k, mid+1, end)
	}
	return searchInRotatedArray(arr, k, start, mid-1)
}

/* 1, 2, 3, 4, 5, 6, 7

2,1,7,6,5,4,3
3 4 5 6 7 1 2

*/
