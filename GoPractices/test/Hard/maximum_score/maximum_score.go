package main

import "fmt"

// Producer and consumer in go
var flag = make(chan bool)

func main() {
	data := make(chan int)
	fmt.Println("staring")
	go producer(data)
	go consumer(data)
	<-flag
	fmt.Println("ending")
}

func producer(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	flag <- true
}

func consumer(data chan int) {
	for true {
		i := <-data
		fmt.Println(i)
	}
}
