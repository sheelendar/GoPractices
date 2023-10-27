package main

import (
	"fmt"
	"sync"
)

type SingleTon struct {
	data string
}

func main() {

	getSingleTonInstance()
	getSingleTonInstance()
}

var singleTon *SingleTon

func getSingleTonInstance() *SingleTon {
	if singleTon != nil {
		fmt.Println("alrady there")
		return singleTon
	}

	if singleTon == nil {
		var lock sync.Mutex
		lock.Lock()
		if singleTon == nil {
			fmt.Println("not there")
			singleTon = &SingleTon{data: "Hi"}
			lock.Unlock()
		} else {
			fmt.Println("alrady there")
		}
	} else {
		fmt.Println("alrady there")
	}
	return singleTon
}
