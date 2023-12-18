/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

Example 1:

Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:

Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:

Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".
*/
package main

import "fmt"

func main() {
	//s := "ab##"
	//	t := "c#d#"
	//s := "a#c"
	//t := "b"

	//s := "a##c"
	//t := "#a#c"

	//s := "xywrrmp"
	//	t := "xywrrmu#p"

	s := "bxj##tw"
	t := "bxo#j##tw"
	fmt.Println(backspaceCompare(s, t))
}
func backspaceCompare(s string, t string) bool {
	i := len(s) - 1
	j := len(t) - 1

	for i >= 0 && j >= 0 {
		count1 := 0
		for i >= 0 && s[i] == '#' {
			count1++
			i--
			if i >= 0 && s[i] != '#' {
				i = i - count1
				count1 = 0
			}
		}
		if count1 > 0 {
			i = i - count1
		}
		count2 := 0
		for j >= 0 && t[j] == '#' {
			j--
			count2++
			if j >= 0 && t[j] != '#' {
				j = j - count2
				count2 = 0
			}
		}
		if count2 > 0 {
			j = j - count2
		}
		if i >= 0 && j >= 0 && s[i] == t[j] {
			i--
			j--
		} else if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}
	}
	if i == j {
		return true
	}
	return false
}
