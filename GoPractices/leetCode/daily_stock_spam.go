package main

func main() {
	obj := Constructor()
	arr := []int{100, 80, 60, 70, 60, 75, 85}
	for i := 0; i < len(arr); i++ {
		param_1 := obj.Next(arr[i])
		print(param_1)
		print(" ")
		print(" ")
	}

}

type StockSpanner struct {
	price []int
}

func Constructor() StockSpanner {
	return StockSpanner{price: make([]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	c := 0
	for p := range this.price {
		if p < price {
			c++
		} else {
			break
		}
	}
	this.price = append(this.price, price)
	return c
}
