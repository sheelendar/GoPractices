package main

import "fmt"

func main() {

	arr := []int{7, 1, 5, 3, 6, 4}
	n := len(arr)
	p := findProfit(arr, n)
	fmt.Print(p)
}

func findProfit(arr []int, n int) int {
	p := 0
	if n == 0 || n == 1 {
		return 0
	}
	temp := arr[0]

	for i := 1; i < n; i++ {
		if temp < arr[i] {
			p = max(p, arr[i]-temp)
		} else {
			temp = arr[i]
		}
	}
	return p
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// {90,110,90,120,80}

//temp:=90
//profit:=20
//90-120 =30
