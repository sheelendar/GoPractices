package main

import (
	"fmt"
	"sort"
)

/*
Twin Strings:
Two strings are called twins if they can be made equivalent by performing some number of operations on one or both strings.
There are two possible operations:

• SwapEven: Swap a character at an even-numbered index with a character at another even-numbered index.
• SwapOdd: Swap a character at an odd-numbered index with a character at another odd-numbered index.

There will be two string arrays. At each index of the two arrays, compare a string from each array,
aligned by index, and store the result in a final string array. The return array should consist of
strings either "Yes" or "No", based on whether the strings compared are twins or not.


Example
firstString = ['abcd", "abcd"]
secondString = ["cbad", "adbc"]

• Compare the two strings firstString[0] and secondString[0].
	One SwapEven operation allows us to swap the characters "a" and "c" of the string "abcd" and make
	it equivalent to "bad" ("abcd" -› "cbad"). firstString[0] and secondString[0] are twins and the answer is "Yes".

• Compare the two strings firstString [1] and secondString[1]. No
	SwapOdd or SwapEven operations can make the two strings abed and adbc equivalent.
	firstString [1] and secondString[1] are not twins and the answer is ["No"].
• The result is c= ["Yes,"No"].

Function Description
Complete the function twins in the editor below.
twins has wo parameters:
string firstString [n]: first array of strings string secondStringin]: second array of strings

Returns:
string[n]: array of strings containing the string "Yes" if firstStringli] and secondStringli] are twins or the string "No" otherwise.
Constraints
• 1 5 n$ 103
• 1 ≤ | firstString[i]l, |secondString[i]| ≤ 100
• firstString[i] and secondStringli] are not guaranteed to have the same length (i.e. | firstString[i]l may not equal secondStringli]|).
• Strings firstString[i] and secondString[i] contain lowercase letters only, in the range asci[a-z].

sample Input 0
STDIN
Function
size n = 3,  firstString = ["cdab", "dcba", "abed"]
size n = 3,  secondString = ["abcd", "abcd","abcdcd"]

Sample Output 0
Yes
No
No
Explanation 0
Given the two string arrays firstString=
*/

func main() {

	/*	firstString := []string{"cdab", "dcba", "abed"}
		secondString := []string{"abcd", "abcd", "abcdcd"}
		fmt.Println(twins(firstString, secondString))
	*/

	/*firstString := []string{"abcd", "abcd"}
	secondString := []string{"cbad", "adbc"}
	fmt.Println(twins(firstString, secondString))
	*/

	firstString := []string{"abcd", "abcd"}
	secondString := []string{"cbad", "adbc"}
	fmt.Println(twins(firstString, secondString))

}

func twins(firstStrings, sedongStrings []string) []string {
	var result []string
	size := len(firstStrings)

	for i := 0; i < size; i++ {
		m := len(firstStrings[i])
		n := len(sedongStrings[i])
		if n != m {
			result = append(result, "No")
		} else {
			str1 := firstStrings[i]
			str2 := sedongStrings[i]
			var fe, fo, se, so string

			for j := 0; j < m; j++ {
				if (j+1)%2 == 1 {
					fo += string(str1[j])
					so += string(str2[j])
				} else {
					fe += string(str1[j])
					se += string(str2[j])
				}
			}
			arr := []byte(fe)
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			fe = string(arr)

			arr = []byte(se)
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			se = string(arr)

			arr = []byte(fo)
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			fo = string(arr)

			arr = []byte(so)
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			so = string(arr)

			if fe == se && fo == so {
				result = append(result, "Yes")
			} else {
				result = append(result, "No")
			}
		}
	}
	return result
}

//fe :=db- > bd
//fo:=ca -> ac
//se:= bd -> bd
//so:=ac -> ac
