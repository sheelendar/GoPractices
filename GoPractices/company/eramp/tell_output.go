package main

import "fmt"

func main() {
	test1()
	test2()
	fmt.Println("...........slice ...........")
	sliceEx()
}

func test1() {
	a := make([]int, 3, 10)
	b := a[:2]
	fmt.Println("a->", a, "b->", b) // 0 0 0,  0 0
	b = append(b, 31)
	fmt.Println(len(a))
	fmt.Println(a, b) //0 0 31, 0 0 31
	a = append(a, 21)
	fmt.Println(a, b) //0 0 31, 0 0 31
	b = append(b, 32)
	fmt.Println(a, b) //0 0 31, 0 0 31,32
	b = append(b, 33)
	fmt.Println(a, b) // 0 0 31,     0 0 31,32,33
}
func test2() {
	var i *int
	var j interface{}
	do(i) //nil
	do(j) //nil
}

func do(i interface{}) {
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func sliceEx() {
	// Creating a slice with a capacity of 100 and a length of 0
	mySlice := make([]int, 0, 100)

	// Length and capacity of the slice
	fmt.Println("Length of mySlice:", len(mySlice))
	fmt.Println("Capacity of mySlice:", cap(mySlice))

	// Filling the slice with data
	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, i*2)
	}

	// Printing the slice
	fmt.Println("mySlice:", mySlice)
}
