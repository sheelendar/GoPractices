/*
Given an n x n matrix and an integer x, find the position of x in the matrix if it is present. Otherwise, print “Element not found”.
Every row and column of the matrix is sorted in increasing order. The designed algorithm should have linear time complexity
Input: mat[4][4] = { {10, 20, 30, 40},  x = 29
                               {15, 25, 35, 45},
                               {27, 29, 37, 48},
                             {32, 33, 39, 50}}
Output: Found at (2, 1)
Explanation: Element at (2,1) is 29
Input : mat[4][4] = { {10, 20, 30, 40},   x = 100
                                {15, 25, 35, 45},
                               {27, 29, 37, 48},
                              {32, 33, 39, 50}};
     	Output: (-1, -1)
*/

package main

import "fmt"

func main() {
	metrix := [][]int{
		{10, 20, 30, 40},
		{15, 25, 35, 45},
		{27, 29, 37, 48},
		{32, 33, 39, 50}}
	k := 51
	i, j := findEelementIndexes(metrix, k)

	fmt.Println(i, ", ", j)
}

func findEelementIndexes(M [][]int, k int) (int, int) {
	m := len(M)
	n := len(M[0])
	i := 0
	j := n - 1 // start from first row and last column
	for i < m && j >= 0 {
		if M[i][j] == k { //  if equal then then i and j
			return i, j
		} else if M[i][j] > k { // if k is less than M[i][j] then discard column
			j--
		} else { //  if k greater than M[i][j] then discard row
			i++
		}
	}
	return -1, -1
}
