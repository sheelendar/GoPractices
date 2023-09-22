package main

import (
	"fmt"
)

func main() {
	word := "maximun"
	fmt.Println(word)
	fmt.Println(reverseVowels(word))
}

func reverseVowels(word string) string {
	arr := []byte(word)
	j := len(arr) - 1
	i := 0
	for i < j {
		if !isVowel(arr[i]) {
			i++
			continue
		}
		if !isVowel(arr[j]) {
			j--
			continue
		}
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
	return string(arr)
}

func isVowel(c byte) bool {
	return (c == 'a' || c == 'A' || c == 'e' ||
		c == 'E' || c == 'i' || c == 'I' ||
		c == 'o' || c == 'O' || c == 'u' ||
		c == 'U')
}
