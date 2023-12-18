package main

import "fmt"

func main() {
	arr := []string{"p", "e", "r", "f", "e", "c", "t", " ", "m", "a", "k", "e", "s", " ", "p", "r", "a", "c", "t", "i", "c", "e"}
	fmt.Print(ReverseWords(arr))
}

func ReverseWords(arr []string) []string {

	size := len(arr)

	if size == 0 {
		return arr
	}

	i := 0
	for k := 0; k < size; k++ {

		if arr[k] == " " {
			reverse(arr, i, k-1)
			i = k + 1
		}
	}
	reverse(arr, i, size-1)
	reverse(arr, 0, size-1)

	return arr
}

func reverse(arr []string, i, j int) {

	for i < j {
		k := arr[i]
		arr[i] = arr[j]
		arr[j] = k
		i++
		j--
	}

}
