package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 2, 2, 3, 4, 7, 8, 8}
	n := len(arr)
	x := 8
	fmt.Println(findFirstIndex(arr, 0, n, x))
	fmt.Print(findLastIndex(arr, 0, n, x))
}

func findLastIndex(arr []int, l, h int, x int) int {
	if l > h {
		return -1
	}
	mid := (l + h) / 2

	if (mid == h-1 || (x < arr[mid+1])) && x == arr[mid] {
		return mid
	} else if arr[mid] <= x {
		return findLastIndex(arr, mid+1, h, x)
	}
	return findLastIndex(arr, l, mid-1, x)
}

func findFirstIndex(arr []int, l, h int, x int) int {
	if l > h {
		return -1
	}
	mid := (l + h) / 2

	if (mid == 0 || x > arr[mid-1]) && x == arr[mid] {
		return mid
	} else if arr[mid] < x {
		return findFirstIndex(arr, mid+1, h, x)
	}
	return findFirstIndex(arr, l, mid-1, x)
}
