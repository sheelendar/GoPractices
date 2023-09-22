package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(happyNumber2(19))
	fmt.Println(happyNumber(20))
}
func happyNumber(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	slow := n
	fast := n
	for true {
		slow = calculateSum(slow)
		fast = calculateSum(fast)
		fast = calculateSum(fast)
		if slow == 1 || fast == 1 {
			return true
		}
		if slow == fast {
			break
		}
	}
	return false
}
func calculateSum(n int) int {
	a := n
	sum := 0
	for a > 0 {
		rem := a % 10
		sum = sum + int(math.Pow(float64(rem), float64(2)))
		a = a / 10
	}
	return sum
}
func happyNumber2(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	a := n
	dp := make(map[int]bool)
	for true {
		sum := 0
		for a > 0 {
			rem := a % 10
			sum = sum + int(math.Pow(float64(rem), float64(2)))
			a = a / 10
		}
		if sum == 1 {
			return true
		}
		if _, ok := dp[sum]; ok {
			return false
		} else {
			dp[sum] = true
		}
		a = sum
	}
	return false
}
