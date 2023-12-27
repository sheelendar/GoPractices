package main

import (
	"fmt"
	"math"
)

/*
You are given k identical eggs and you have access to a building with n floors labeled from 1 to n.

You know that there exists a floor f where 0 <= f <= n such that any egg dropped at a floor higher than f will break,

	and any egg dropped at or below floor f will not break.

Each move, you may take an unbroken egg and drop it from any floor x (where 1 <= x <= n).
If the egg breaks, you can no longer use it. However, if the egg does not break, you may reuse it in future moves.

Return the minimum number of moves that you need to determine with certainty what the value of f is.

Example 1:
Input: k = 1, n = 2
Output: 2
Explanation:
Drop the egg from floor 1. If it breaks, we know that f = 0.
Otherwise, drop the egg from floor 2. If it breaks, we know that f = 1.
If it does not break, then we know f = 2.
Hence, we need at minimum 2 moves to determine with certainty what the value of f is.

Example 2:
Input: k = 2, n = 6
Output: 3

Example 3:
Input: k = 3, n = 14
Output: 4
*/
func main() {
	k := 3
	n := 14
	fmt.Println(superEggDrop(k, n))
}

func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		dp[1][i] = i
	}

	for i := 2; i <= k; i++ {
		for j := 1; j <= n; j++ {
			res := math.MaxInt
			left := 1
			right := j
			for left <= right {
				mid := (left + right) / 2
				a := dp[i-1][mid-1]
				b := dp[i][j-mid]
				max := Max(a, b) + 1
				res = Min(max, res)
				if a == b {
					break
				} else if a > b {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
			dp[i][j] = res
		}
	}
	return dp[k][n]
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* Recursion Time Out
func helper(k int , n int,dp map[string]int) int {
    if k==1 || n<=1{
        return n
    }

    key:=fmt.Sprintf("%d_%d",k,n);
    if v,ok:=dp[key]; ok{
        return v
    }

    minAttempt:=math.MaxInt
    for x:=1; x<=n; x++{
        res:=Max(helper(k-1,x-1,dp), helper(k,n-x,dp)) + 1 // adding 1 for current attempt and Max of remaining attempt

        if res < minAttempt{
            minAttempt=res
        }
    }
    dp[key]=minAttempt
    return minAttempt
}

func Max(a,b int)int {
    if a>b{
        return a
    }
    return b
}
*/
