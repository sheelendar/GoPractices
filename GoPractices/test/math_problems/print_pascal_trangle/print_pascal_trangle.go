package main

import "fmt"

/*

Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle. In Pascal's triangle, each number is the sum of the two numbers directly above it as shown: https://leetcode.com/problems/pascals-triangle-ii/
Example 1: Input: rowIndex = 3  :   Output: [1,3,3,1]
Example 2:   Input: rowIndex = 0  :    Output: [1]
Example 3:   Input: rowIndex = 1  :  Output: [1,1]
Ans: 		public static List<Integer> getRow2(int rowIndex) {
	List<Integer> ret = new ArrayList<Integer>();
	ret.add(1);
	for (int i = 1; i <= rowIndex; i++) {
		for (int j = i ; j >= 1; j--) {
			int tmp = ret.get(j - 1) + ret.get(j);
			ret.set(j, tmp);
		}
		ret.add(1);
	}
	return ret;
}

*/

func main() {
	n := 5
	fmt.Println(getRow(n))

}
func getRow(n int) []int {
	res := make([]int, n+1)
	res[0] = 1
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			res[j] = res[j] + res[j-1]
		}
	}
	return res
}
