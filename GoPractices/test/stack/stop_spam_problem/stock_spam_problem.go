package main

import "fmt"

/*
The stock span problem is a financial problem where we have a series of N daily price quotes for a stock and we need to calculate
the span of the stock’s price for all N days.The span Si of the stock’s price on a given day i is defined as the maximum number
of consecutive days just before the given day,for which the price of the stock on the current day is less than its price on the given day.
Examples:(https://www.geeksforgeeks.org/the-stock-span-problem/)

	Input: N = 7, price[] = [100 80 60 70 60 75 85]
	Output: 1 1 1 2 1 4 6
	Explanation: Traversing the given input span for 100 will be 1, 80 is smaller than 100 so the span is 1, 60 is smaller than 80 so the span is 1, 70 is greater than 60 so the span is 2 and so on. Hence the output will be 1 1 1 2 1 4 6.
*/

func main() {
	price := []int64{100, 80, 60, 70, 60, 75, 85} // output {1 1 1 2 1 4 6}
	//price := []int64{10, 4, 5, 90, 120, 80} // output: {1 1 2 4 5 1}
	size := len(price)
	fmt.Println(calculateSpam(price, size))
}
func calculateSpam(price []int64, size int) []int {
	if size == 0 {
		return nil
	}
	var result []int
	result = append(result, 1)
	var s []int
	s = append(s, 0)
	for i := 1; i < size; i++ {
		for len(s) > 0 && price[s[len(s)-1]] < price[i] {
			s = s[:len(s)-1]
		}
		if len(s) > 0 {
			result = append(result, i-s[len(s)-1])
		} else {
			result = append(result, i+1)
		}
		s = append(s, i)
	}
	return result
}
