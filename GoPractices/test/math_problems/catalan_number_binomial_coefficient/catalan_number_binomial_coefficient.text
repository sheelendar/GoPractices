package main

func main() {
	n := 10
	for i := 0; i < n; i++ {
		println(catalanNumber(i))
	}
}

func catalanNumber(n int) int {
	c := binomialCofficient(2*n, n)
	return c / (n + 1)
}

func binomialCofficient(n, k int) int {
	if n == 0 || n == 1 {
		return 1
	}
	res := 1
	if n-k < k {
		k = n - k
	}
	for i := 0; i < k; i++ {
		res = res * (n - i)
		res = res / (i + 1)
	}
	return res
}
