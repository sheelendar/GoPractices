package main

import "fmt"

/*
A celebrity is a person who is known to all but does not know anyone at a party. If you go to a party of N people,
find if there is a celebrity in the party or not.
A square NxN matrix M[][] is used to represent people at the party such that if an element of row i and column j
is set to 1 it means ith person knows jth person. Here M[i][i] will always be 0.
Note: Follow 0 based indexing.
Follow Up: Can you optimize it to O(N)

Example 1:

Input:
N = 3
M[][] = {{0 1 0},

	{0 0 0},
	{0 1 0}}

Output: 1
Explanation: 0th and 2nd person both
know 1. Therefore, 1 is the celebrity.
*/
func main() {
	M := [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 1, 0}}
	// Output: element is celebrity 1
	fmt.Println(findCelebrity(M))
	fmt.Println(Celebrity(M))
}
func Celebrity(M [][]int) int {
	n := len(M)
	if n == 1 {
		return -1
	}
	i := 0
	j := n - 1
	for i < j {
		if M[j][i] == 1 { // if j knows i then j can not be celerity
			j--
		} else { // j doesn't know i so i cant be celebrity
			i++
		}
	}
	cele := i
	for i := 0; i < n; i++ {
		if M[cele][i] == 1 {
			return -1
		}
		if M[i][cele] != 1 && cele != i {
			return -1
		}
	}
	return cele
}

func findCelebrity(M [][]int) int {
	n := len(M)
	if n == 1 {
		return -1
	}
	var stack []int
	for i := 0; i < n; i++ {
		stack = append(stack, i)
	}
	for len(stack) > 1 {
		size := len(stack) - 1
		i := stack[size]
		stack = stack[:size]
		j := stack[size-1]
		stack = stack[:size-1]
		if M[i][j] != 1 && M[j][i] == 1 {
			stack = append(stack, i)
		}
		if M[j][i] != 1 && M[i][j] == 1 {
			stack = append(stack, j)
		}
	}
	if len(stack) <= 0 {
		return -1
	}
	cele := stack[0]
	for i := 0; i < n; i++ {
		if M[cele][i] == 1 {
			return -1
		}
		if M[i][cele] != 1 && cele != i {
			return -1
		}
	}
	return cele
}
