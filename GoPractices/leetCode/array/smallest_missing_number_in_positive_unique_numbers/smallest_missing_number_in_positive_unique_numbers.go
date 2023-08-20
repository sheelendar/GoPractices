package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 1, 0, 5}
	fmt.Print(GetDifferentNumber(arr))
}

func GetDifferentNumber(arr []int) int {
	n := len(arr)

	if n == 0 {
		return 0
	}
	if n == 1 && arr[0] != 0 {
		return 0
	}
	temp := 0
	for i := 0; i < n; i++ { // arr = [2, 1, 0, 5]
		j := abs(arr[i])
		if j < n {
			if arr[j] == 0 {
				temp = arr[i]
			}
			arr[j] = -arr[j]
		}

	}
	m := n
	for i := 0; i < n; i++ {
		if arr[i] > 0 && arr[i] != 0 {
			m = i
			break
		}
	}
	if m == temp {
		return n
	}
	return m

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
