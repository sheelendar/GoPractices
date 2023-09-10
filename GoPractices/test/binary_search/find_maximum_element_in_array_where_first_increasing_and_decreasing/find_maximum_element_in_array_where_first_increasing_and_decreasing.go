package main

/*
Find the maximum element in an array which is first increasing and then decreasing
Given an array of integers which is initially increasing and then decreasing, find the maximum value in the array.
https://www.geeksforgeeks.org/find-the-maximum-element-in-an-array-which-is-first-increasing-and-then-decreasing/
Examples :

Input: arr[] = {8, 10, 20, 80, 100, 200, 400, 500, 3, 2, 1}
Output: 500

Input: arr[] = {1, 3, 50, 10, 9, 7, 6}
Output: 50

Corner case (No decreasing part)
Input: arr[] = {10, 20, 30, 40, 50}
Output: 50

Corner case (No increasing part)
Input: arr[] = {120, 100, 80, 20, 0}
Output: 120
*/
func main() {
	arr := []int64{8, 10, 20, 80, 100, 200, 400, 500, 3, 2, 1}

	println("mexElement: ", findMexElement(arr, 0, len(arr)-1))
}

func findMexElement(arr []int64, l, h int) int64 {
	if l == h {
		return arr[l]
	}
	if l+1 == h && arr[l] >= arr[h] {
		return arr[l]
	}
	if l+1 == h && arr[l] < arr[h] {
		return arr[h]
	}
	mid := (l + h) / 2
	if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
		return arr[mid]
	}
	if arr[mid-1] > arr[mid] && arr[mid] > arr[mid+1] {
		return findMexElement(arr, l, mid-1)
	} else {
		return findMexElement(arr, mid+1, h)
	}
}
