package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Write a golang program which spins up 3 go-routines. The first go-routine is responsible for generating a list of 4
random numbers every 2 seconds where the last number represents the “value” of the list. The second go-routine is
responsible for generating a list of 21 numbers every 3 seconds where the last number represents the “value” of the list.
 The third go-routine is responsible for receiving these 2 lists and concatenating these 2 lists and provides the highest
  value list to the main program. The value of the combined list is the sum of the values of both lists.

Follow up:

Instead of one go-routine receiving and concatenating a list, build a system where there are 3 more go-routines
 receiving and concatenating these lists. The fastest one is delivered to the main program.

L1: 1 2 3 4

L2: 1 2 3 4 ….. 21

L1 + L2: 1 2 3 4 1 2 3 4 …. 21 (Value = 4 + 21 = 25)

L1: 1 2 2 1

L2: 1 2 3 4 ….. 2

L1 + L2: 1 2 2 1 1 2 3 4 ….. 2 (Value = 1 + 2 = 3)

L1: 1 2 3 10

L2: 1 2 3 4 ….. 23

L1 + L2: 1 2 3 10 1 2 3 4 ….. 23 (Value = 10 + 23 = 33)

L1: 1 2 3 3

L2: 1 2 3 4 …. 27

L1 + L2: 1 2 3 3 1 2 3 4 …. 27 (Value = 3 + 27 = 30)
*/

func main() {

	var wg sync.WaitGroup
	ch1 := make(chan []int)
	ch2 := make(chan []int)
	n1 := 4
	n2 := 21
	wg.Add(1)
	go genrateList1(ch1, &wg, n1)

	wg.Add(1)
	go genrateList2(ch2, &wg, n2)

	wg.Add(1)
	go maxNumberSum(ch1, ch2, &wg)

	wg.Wait()

}

func genrateList1(ch chan []int, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	count := 0
	var list []int
	for count < n {
		list = append(list, rand.Intn(n+1))
		count++
	}
	ch <- list
	close(ch)
}

func genrateList2(ch chan []int, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	count := 0
	var list []int
	for count < n {
		list = append(list, rand.Intn(n+1))
		count++
	}
	ch <- list
	close(ch)
}
func maxNumberSum(ch1, ch2 chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var list1 []int
	var list2 []int

	select {
	case data := <-ch1:
		list1 = data
	}

	select {
	case data := <-ch2:
		list2 = data
	}

	fmt.Println(list1)
	fmt.Println(list2)
	value := list1[len(list1)-1] + list2[len(list2)-1]
	fmt.Println(value)
}
