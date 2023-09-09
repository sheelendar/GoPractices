package main

import "fmt"

func main() {
	fmt.Println("Total binary tree: ", TotalNumberOFTree(5))
	fmt.Println("Total Bst: ", TotalNumberOFBST(5))
	fmt.Println("binomial cofficient: ", binomialCofficient(5, 2))
}

func TotalNumberOFTree(n int) int {
	return catelan(n) * factorial(n)
}
func TotalNumberOFBST(n int) int {
	return catelan(n)
}

func catelan(n int) int {
	bico := binomialCofficient(2*n, n)
	return bico / (n + 1)
}
func binomialCofficient(n, k int) int {
	if k > n-k {
		k = n - k
	}
	res := 1
	for i := 0; i < k; i++ {
		res = res * (n - i)
		res = res / (i + 1)
	}
	return res
}
func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	return res
}
