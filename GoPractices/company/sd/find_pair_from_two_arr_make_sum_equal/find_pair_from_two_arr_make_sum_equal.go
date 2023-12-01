/*A = [4,1,2,1,1,2]
B = [3,3,6,3]

Output -
(1,3)

(x,y) => A,B such that if you swap x and y, sum of arrays A,B become equal.
*/

package main

import "fmt"

func main() {
	A := []int{4, 1, 2, 1, 1, 2} //sum1=11
	B := []int{3, 3, 6, 3}       //sum1=15
	m := len(A)
	n := len(B)
	sum1 := 0
	sum2 := 0
	for i := 0; i < m; i++ {
		sum1 = sum1 + A[i]
	}
	for i := 0; i < n; i++ {
		sum2 = sum2 + B[i]
	}
	swap(sum1, sum2, m, n, A, B)
}

func swap(sum1, sum2, m, n int, A, B []int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			s1 := sum1 - A[i] + B[j]
			s2 := sum2 + A[i] - B[j]

			if s1 == s2 {
				fmt.Println(A[i], B[j])
				fmt.Println(s1, s2)
				return
			}
		}
	}
}
