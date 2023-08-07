package main

import "fmt"

/*
Busiest Time in The Mall

The Westfield Mall management is trying to figure out what the busiest moment at the mall was last year. You’re given data extracted from the mall’s door detectors. Each data point is represented as an integer array whose size is 3. The values at indices 0, 1 and 2 are the timestamp, the count of visitors, and whether the visitors entered or exited the mall (0 for exit and 1 for entrance), respectively. Here’s an example of a data point: [ 1440084737, 4, 0 ].

Note that time is given in a Unix format called Epoch, which is a nonnegative integer holding the number of seconds that have elapsed since 00:00:00 UTC, Thursday, 1 January 1970.

Given an array, data, of data points, write a function findBusiestPeriod that returns the time at which the mall reached its busiest moment last year. The return value is the timestamp, e.data_stractures. 1480640292. Note that if there is more than one period with the same visitor peak, return the earliest one.

Assume that the array data is sorted in an ascending order by the timestamp. Explain your solution and analyze its time and space complexities.

Example:

input:  data = [ [1487799425, 14, 1],

	[1487799425, 4,  0],
	[1487799425, 2,  0],
	[1487800378, 10, 1],
	[1487801478, 18, 0],
	[1487801478, 18, 1],
	[1487901013, 1,  0],
	[1487901211, 7,  1],
	[1487901211, 7,  0] ]

output: 1487800378 # since the increase in the number of people

	# in the mall is the highest at that point

Constraints:
*/

func main() {
	data := [][]int{{1487799425, 14, 1},
		{1487799425, 4, 0},
		{1487799425, 2, 0},
		{1487800378, 10, 1},
		{1487801478, 18, 0},
		{1487801478, 18, 1},
		{1487901013, 1, 0},
		{1487901211, 7, 1},
		{1487901211, 7, 0}}
	fmt.Print("Busiest day ", FindBusiestPeriod(data))
}

func FindBusiestPeriod(data [][]int) int {
	max_count := 0
	size := len(data)
	sum := 0
	maxIndex := -1
	for i := 0; i < size; i++ {

		if data[i][2] == 1 {
			sum = sum + data[i][1]

		} else if data[i][2] == 0 {
			sum = sum - data[i][1]
		}
		if i < size-1 && data[i][0] == data[i+1][0] {
			continue
		}
		if sum > max_count {
			sum = max_count
			maxIndex = data[i][0]
		}
	}
	return maxIndex
}
