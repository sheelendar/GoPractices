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
	// if array sorted and rotated then either (arr[start] <= arr[mid]) will true or arr[mid] <= arr[high] will be true
	//  checking arr[start] <= arr[mid] is true then check value exits in sorted array
	if arr[start] <= arr[mid] {
		// check value exist in then search here.
		if k >= arr[start] && k <= arr[mid] {
			return searchInRotatedArray(arr, k, start, mid-1)
		}
		// otherwise search in unsorted part
		return searchInRotatedArray(arr, k, mid+1, end)
	}

	// considering here arr[start] <= arr[mid] is false then  arr[mid] <= arr[high]  is true
	if k >= arr[mid] && k <= arr[end] {
		return searchInRotatedArray(arr, k, mid+1, end)
	}
	return searchInRotatedArray(arr, k, start, mid-1)
}

/* 1, 2, 3, 4, 5, 6, 7

2,1,7,6,5,4,3
3 4 5 6 7 1 2

*/
