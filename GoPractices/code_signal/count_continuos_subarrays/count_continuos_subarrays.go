/*
	You are given an array of integers numbers and a positive integer k Your task is to count the number of
	contiguous subarrays within numbers that contains at least k pairs of elements with duplicate values.

More formally, count the number of contiguous subarrays numbers li..jl(is g ) for which there are
* * elements (with pairwise distinct indices i ≤ i1, 31, 12, j2, ..., ik, jk S § ) with each
pair having the same value - numbers [ill = numbers [jil, numbers [i2l = numbers[j21, ...
numbers [ik] = numbers [jk)

• For numbers = [0, 1, 0, 1, 0] and k = 2, the output should be solution(numbers,2)= 3
There are 3 subarrays that satisfy the criteria of containing at least k = 2 pairs of duplicate
values:
o
numbers [0..3] = [0, 1, 0, 1] with numbers [0] = 0 = numbers [2] and numbers (1]= 1 = numbers [3]
numbers [1..4] = [1, 0, 1, 0] with numbers [1] = 1 = numbers [31 and numbers [2] 0 = numbers [4]
• numbers [0..4] = [0, 1, 0, 1, 0] with numbers [0] = 0 = numbers [2] ,numbers (1] = 1 = numbers [3] , and numbers [21 = 0 = numbers [41
In each of these subarrays, there is at least one pair of o s and one pair of 1 s.

For numbers = [2, 2, 2, 2, 2, 2] and k = 3, the output should be solution (numbers, k) = 1.
numbers [0..5] = [2, 2, 2, 2, 2, 2]

For numbers = [1, 3, 3, 1] and k = 1, the output should be solution (numbers, k) =4
There are 4 subarrays that satisfy the criteria of containing at least k = 1 pair of duplicate

numbers [0..2] = [1, 3, 3]
numbers [0..3] = [1, 3, 3, 1]
numbers [1..2] = [3, 3]
numbers [1..3] = [3, 3, 1)

values:
*/
package main

func main() {

}
