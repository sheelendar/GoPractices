package main

import "fmt"

/**
 * Timmy the sheep has devised a strategy that helps her fall asleep faster.
 First, she picks a number N. Then she starts naming N, 2 × N, 3 × N, and so on.
  Whenever she names a number, she thinks about all of the digits in that number.
  She keeps track of which digits (0, 1, 2, 3, 4, 5, 6, 7, 8, and 9)
  she has seen at least once so far as part of any number she has named.
  Once she has seen each of the ten digits at least once, she will fall asleep.
Timmy must start with N and must always name (i + 1) × N directly after i × N.
For example, suppose that Timmy picks N = 1692. She would count as follows:

N = 1692. Now she has seen the digits 1, 2, 6, and 9.
2N = 3384. Now she has seen the digits 1, 2, 3, 4, 6, 8, and 9.
3N = 5076. Now she has seen all ten digits, and falls asleep.
What is the last number that she will name before falling asleep? If she will count forever, print INSOMNIA instead.
*/
/*
n=1692
dp[]
a =1692
i=1
ans=1692
dp[2]=true,  a =169
dp[9]=true ,a =16
dp[6]=true ,a =1
dp[1]=true ,a =0

i=2
a=3384
ans=3384

dp[4]=true, a =338
dp[8]=true ,a =33
dp[3]=true ,a =3
dp[1]=true ,a =0


 1 ->  all digites
 2 -> 2,4,6,8,10, ->all digites
 3 -> 3,6,9,12,15 - > all digites
 99 -> 1,8,6 ,5,9,7

*/
func findLastNum(n int) int {
	uniqueDigites := make(map[int]bool)
	if n == 0 {
		return -1
	}
	i := 1
	ans := -1
	for true {
		nextNumber := i * n
		ans = nextNumber
		for nextNumber > 0 { // 1692/10 = 169, 169/10= 16
			rem := nextNumber % 10       // take last digite from end
			nextNumber = nextNumber / 10 // remaining digites after removing last one
			if !uniqueDigites[rem] {     // check if uniqueDigites don't have digit then put it into map
				uniqueDigites[rem] = true
			}
			if len(uniqueDigites) == 10 { // check if uniqueDigites size is 10 then we got all digites from
				return ans
			}
		}
		i++
	}
	return ans
}

func main() {
	fmt.Println(findLastNum(1692))
	/*for i := 0; i < 1000; i++ {
		fmt.Println(findLastNum(i))
	}*/
}
