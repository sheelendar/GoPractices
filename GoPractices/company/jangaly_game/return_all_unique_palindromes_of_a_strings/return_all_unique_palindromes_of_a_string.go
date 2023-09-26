package main

import "fmt"

/*
Return all unique palindromes of a string equal to the string length if not possible then return "".
https://www.geeksforgeeks.org/print-palindromic-permutations-given-string-alphabetic-order/
*/

func main() {
	s := "malayalam"
	res := palindromesOfLengthGivenString(s)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func palindromesOfLengthGivenString(s string) []string {
	freMap := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		freMap[s[i]]++
	}
	totalOddLatter := 0
	oddLatter := byte(0)
	totalLength := 0
	for latter, count := range freMap {
		if count%2 != 0 {
			totalOddLatter++
			oddLatter = latter
		}
		freMap[latter] = count / 2
		totalLength = totalLength + count/2

	}
	if totalOddLatter > 2 {
		return []string{}
	}
	var result []string
	findPalindromes("", freMap, 0, totalLength, oddLatter, &result)
	return result
}

func findPalindromes(palin string, freMap map[byte]int, l, r int, latter byte, result *[]string) {
	if l == r {
		rev := reverse(palin)
		res := ""
		if latter == 0 {
			res = fmt.Sprintf("%s%s", palin, rev)
		} else {
			res = fmt.Sprintf("%s%s%s", palin, string(latter), rev)
		}
		*result = append(*result, res)
	} else {
		for char, count := range freMap {
			if count > 0 {
				freMap[char] = count - 1
				findPalindromes(palin+string(char), freMap, l+1, r, latter, result)
				freMap[char] = count
			}
		}
	}
}

func reverse(palin string) string {
	n := len(palin)
	res := []byte(palin)
	i := 0
	j := n - 1
	for i < j {
		temp := res[i]
		res[i] = res[j]
		res[j] = temp
		i++
		j--
	}
	return string(res)
}
