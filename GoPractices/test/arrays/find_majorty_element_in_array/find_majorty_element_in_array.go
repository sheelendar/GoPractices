package main

import "fmt"

/*
Find the majority element in the array. A majority element in an array A[] of size n is an element that appears more
than n/2 times (and hence there is at most one such element).
Examples :

	Input : A[]={3, 3, 4, 2, 4, 4, 2, 4, 4}
	Output : 4
	Explanation: The frequency of 4 is 5 which is greater than the half of the size of the array size.

	Input : A[] = {3, 3, 4, 2, 4, 4, 2, 4}
	Output : No Majority Element
	Explanation: There is no element whose frequency is greater than the half of the size of the array size.
*/

func main() {
	//arr := []int{3, 3, 4, 2, 4, 4, 2, 4, 4} //true
	arr := []int{3, 3, 4, 2, 4, 4, 2, 4} // false
	n := len(arr)
	a := checkIfAnyElementCanMajority(arr, n)
	fmt.Println(isElementMajority(arr, n, a))
}

func isElementMajority(arr []int, n int, a int) bool {
	count := 0
	for i := 0; i < n; i++ {
		if a == arr[i] {
			count++
		}
	}
	if count > n/2 {
		return true
	}
	return false
}

func checkIfAnyElementCanMajority(arr []int, n int) int {
	index := 0
	count := 1
	for i := 1; i < n; i++ {
		if arr[index] == arr[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			index = i
			count = 1
		}
	}
	return arr[index]
}
