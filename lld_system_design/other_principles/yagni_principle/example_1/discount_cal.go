package main

import "fmt"

// Simple and necessary functionality only
func calculateDiscount(price float64, discount float64) float64 {
	return price * (1 - discount/100)
}

func main() {
	price := 100.0
	fmt.Println("Discounted Price:", calculateDiscount(price, 10))
}
