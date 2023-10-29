/*
Consider a rectangular h × w board with all cells initially white. You are to process several queries of the following types:

    "x a b" - color the white cell (a, b) (0-based coordinates, the first one is a row index, and the second one is a column index) black;
    "> a b" - find the leftmost white cell to the right of the white cell (a, b);
    "< a b" - find the rightmost white cell to the left of the white cell (a, b);
    "v a b" - the same, but the search should be done downwards;
    "^ a b" - the same, but the search should be done upwards;

For each query, except the ones of the first type, find the answer.

Example

For h = 3, w = 5, and
queries = ["v 1 2", "x 2 2", "v 1 2", "> 2 1", "x 2 3", "> 2 1", "< 2 0"],
the output should be
solution(h, w, queries) = [[2, 2], [-1, -1], [2, 3], [2, 4], [-1, -1]].

Check out the image above to see the state of the board after each query of the first type:

Input/Output

    [execution time limit] 3 seconds (java)

    [memory limit] 1 GB

    [input] integer h

    A positive integer.

    Guaranteed constraints:
    1 ≤ h ≤ 500.

    [input] integer w

    A positive integer.

    Guaranteed constraints:
    1 ≤ w ≤ 500.

    [input] array.string queries

    Queries in the above-described format.

    Guaranteed constraints:
    5 ≤ queries.length ≤ 104.

    [output] array.array.integer

    For each query except the ones of the first type, store the answer's coordinates in the array. If the desired cell doesn't exist, store [-1, -1] instead. The answers should be stored in the same order as the queries.

[Java] Syntax Tips

// Prints help message to the console
// Returns a string
//
// Globals declared here will cause a compilation error,
// declare variables inside the function instead!
String helloWorld(String name) {
    System.out.println("This prints to the console when you Run Tests");
    return "Hello, " + name;
}

1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
Tests
Custom Tests
Test 1
Input
Return Value
Console Output
Error Output

h: 3
w: 5
queries:
["v 1 2",
 "x 2 2",
 "v 1 2",
 "> 2 1",
 "x 2 3",
 "> 2 1",
 "< 2 0"]

Test 2
Test 3
Test 4
Test 5
Test 6
Test 7
Test 8
Test 9

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	h := 3
	w := 5
	queries := []string{"v 1 2", "x 2 2", "v 1 2", "> 2 1", "x 2 3", "> 2 1", "< 2 0"}

	//solution(h, w, queries) = [[2, 2], [-1, -1], [2, 3], [2, 4], [-1, -1]].
	fmt.Println(solution(h, w, queries))
}
func solution(h int, w int, queries []string) [][]int {
	dp := make(map[byte]int)
	dp['x'] = 0
	dp['>'] = 1
	dp['<'] = 2
	dp['v'] = 3
	dp['^'] = 4
	board := make([][]int, h)
	for i := 0; i < h; i++ {
		board[i] = make([]int, w)
	}

	var result [][]int
	for k := 0; k < len(queries); k++ {
		arr := strings.Split(queries[k], " ")
		i, _ := strconv.Atoi(arr[1])
		j, _ := strconv.Atoi(arr[2])

		if arr[0] == "x" {
			board[i][j] = 1
		} else {
			res := DFS(board, i, j, h, w, arr[0])
			result = append(result, res)
		}

	}
	return result
}
func DFS(board [][]int, row, col, m, n int, symbol string) []int {

	res := []int{-1, -1}
	if symbol == ">" {
		for j := col + 1; j < n; j++ {
			if board[row][j] == 0 {
				return []int{row, j}
			}
		}

	} else if symbol == "<" {
		for j := col - 1; j >= 0; j-- {
			if board[row][j] == 0 {
				return []int{row, j}
			}
		}
	} else if symbol == "v" {
		for i := row + 1; i < m; i++ {
			if board[i][col] == 0 {
				return []int{i, col}
			}
		}
	} else if symbol == "^" {
		for i := row - 1; i >= 0; i-- {
			if board[i][col] == 0 {
				return []int{i, col}
			}
		}
	}
	return res
}
