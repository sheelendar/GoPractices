package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ABC"
	printAllPermutations(str, 0, len(str))
}

func printAllPermutations(str string, l, r int) {
	if l == r {
		fmt.Println(str)
	} else {
		for i := l; i < r; i++ {
			str = swapValues(str, l, i)
			printAllPermutations(str, l+1, r)
			str = swapValues(str, l, i)
		}
	}
}

func swapValues(str string, l int, i int) string {
	stringArr := []byte(str)
	temp := stringArr[l]
	stringArr[l] = stringArr[i]
	stringArr[i] = temp
	return string(stringArr)
}

func dd() {
	strings.Contains("", "substr")
	strings.LastIndex()
	strings.I
}
