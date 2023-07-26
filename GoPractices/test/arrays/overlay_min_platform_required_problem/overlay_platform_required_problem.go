package main

import (
	"fmt"
	"sort"
)

/*
Given the arrival and departure times of all trains that reach a railway station, the task is to find the minimum number
of platforms required for the railway station so that no train waits. We are given two arrays that represent the arrival
and departure times of trains that stop.
Examples: (https://www.geeksforgeeks.org/minimum-number-platforms-required-railwaybus-station/)

	Input: arr[] = {9:00, 9:40, 9:50, 11:00, 15:00, 18:00}, dep[] = {9:10, 12:00, 11:20, 11:30, 19:00, 20:00}
	Output: 3
	Explanation: There are at-most three trains at a time (time between 9:40 to 12:00)

	Input: arr[] = {9:00, 9:40}, dep[] = {9:10, 12:00}
	Output: 1
	Explanation: Only one platform is needed.
*/

func main() {
	arr := []int{900, 940, 950, 1100, 1500, 1800}
	dep := []int{910, 1200, 1120, 1130, 1900, 2000}
	fmt.Print(requiredPlatfomr(arr, dep, len(arr)))
}

func requiredPlatfomr(arr []int, dep []int, n int) int {
	requirdPF := 0
	maxRequiredPF := 0
	sort.Ints(arr)
	sort.Ints(dep)
	j := 0
	i := 0
	for i < n && j < n {
		if arr[i] <= dep[j] {
			requirdPF++
			i++
		} else if arr[i] > dep[j] {
			requirdPF--
			j++
		}
		if requirdPF > maxRequiredPF {
			maxRequiredPF = requirdPF
		}
	}
	return maxRequiredPF
}
