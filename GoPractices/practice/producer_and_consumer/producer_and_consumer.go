package main

import (
	"fmt"
	"sync"
)

var flag = make(chan bool)

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
func main() {
	data := make(chan int)
	fmt.Println("staring")
	go producer(data)
	go consumer(data)
	<-flag
	fmt.Println("ending")

	fmt.Println()
	fmt.Println(".....................second example of producer and consumer with waitGroup.......................")
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	go createNumberProducer(ch, &wg)
	wg.Add(1)
	go dislayNumberConsumer(ch, &wg)
	wg.Wait()

}

func dislayNumberConsumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for true {
		data, ok := <-ch //getting data from channel
		fmt.Println("pulling ", data)
		if !ok {
			break
		}
	}
}
func createNumberProducer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("pushing ", i)
		ch <- i // puting data into channel
	}
	close(ch)
}
