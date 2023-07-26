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

type Stack struct {
	stack []int64
}

func (s *Stack) Push(data int64) {
	s.stack = append(s.stack, data)
}
func (s *Stack) Pop() int64 {
	n := len(s.stack)
	if s.stack == nil {
		return -1
	}
	if n == 1 {
		data := s.stack[n-1]
		s.stack = nil
		return data
	}
	data := s.stack[n-1]
	s.stack = s.stack[0 : n-1]
	return data

}
func (s *Stack) IsEmpty() bool {
	n := len(s.stack)
	if n <= 0 {
		return true
	}
	return false
}
func (s *Stack) Top() int64 {
	n := len(s.stack)
	if n > 0 {
		return s.stack[n-1]
	}
	return -1
}

func main() {
	price := []int64{100, 80, 60, 70, 60, 75, 85}
	size := len(price)
	result := make([]int, size)
	calculateSpam(price, size, result)
}

func calculateSpam(price []int64, size int, result []int) {
	if size == 0 {
		return
	}
	for i := 0; i < size; i++ {
		result[i] = 1
	}
	s := Stack{}
	s.Push(int64(size - 1))

	for i := size - 2; i >= 0; i-- {
		if price[int(s.Top())] > price[i] {
			s.Push(int64(i))
		} else {
			for !s.IsEmpty() && price[int(s.Top())] <= price[i] {
				index := s.Pop()
				result[index] = int(index) - i
			}
			s.Push(int64(i))
		}
	}
	for i := 0; i < size; i++ {
		fmt.Print(result[i])
		fmt.Print(" ")
	}
}
