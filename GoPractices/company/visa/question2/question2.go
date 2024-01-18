package main

/*
You are given an array of integers memory consisting of o sand 1 s which indicates whether the
corresponding memory unit is free or not. memory [i] = o means that the ith memory unit is free,
and memory [i] = 1 means it's occupied.
The memory is aligned with segments of & units so all occupied memory blocks must start at an index
divisible by 8 (e.g. 0, 8, 16, etc).
Your task is to perform two types of queries:
• alloc X: Find the left-most aligned memory block of × consecutive free memory units and mark these units as occupied (ie: find the left-most contiguous subarray of o s, starting at the position start which is divisible by 8 , and replace all these memory units with 1 s).
• If there is no proper aligned memory block with × consecutive free units, return -1;
otherwise return the index of the first position of the allocated block segment and assign an
ID to every single element in the block, based on an atomic counter (the counter starts at 1 and is incremented on every successtul alloc operation).
• Note: × may be greater than 8, so the block may cover more than one memory segment.
• erase ID: If there exists an allocated memory block with element ids equal to ID , free all its
memory units (set all of the bits in the block to 0 ).
Return the length of the deleted memory block. If there is no such ID or the block with this
ID has already been deleted, return -1
The queries are given in the following format:
• queries is an array of 2-elements arrays;
• if queries [i] [0] = 0 then this is an alloc type query, where x = queries [i] [l];
• if queries [i] [0] = 1 then this is an erase type query, where ID = queries [i] [1]

memory: [0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1]
queries: [0,2][0, 1], [0, 1], [1, 1], [0, 3], [1, 4], [0, 4],
the output should be solution(memory queries) = [8, 0, -1, 2, 8, -1, -1]

Expand to see the example video.
[0, 2] corresponds to alloc 2, which allocates a memory block from units 8 to 9, as
8 is the first aligned position with 2 free slots. ID = 1 is also assigned to this segment.
After this operation memory = [0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1].
Return the memory block starting index, 8 .

• [0, 1] corresponds to alloc 1, which allocates a memory block from units o to 1.
After this operation, memory = [1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1] .
ID = 2 is also assigned to this segment.

• [0, 1] corresponds to alloc 1, there is not empty memory block .
ans:  -1

• [1, 1] corresponds to erase 1, which means need to remove memory with ID 1 and we free 2 spaces.
After this operation, memory = [1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1] .
ans: 2 free memory spaces

• [0, 3] corresponds to alloc 2, which allocates a memory block from units 8 to 10, as.
After this operation, memory = [1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1] .
ID = 3 is also assigned to this segment.
ans:=8

• [1, 4] corresponds to erase 4 ID, which means need to remove memory with ID 4 there is not ID with 4.
ans: -1

• [0, 4] corresponds to alloc 4, there is not empty memory block .
ans:  -1
*/
func main() {
	memory := []int{0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	queries := [][]int{{0, 2}, {0, 1}, {0, 1}, {1, 1}, {0, 3}, {1, 4}, {0, 4}}
	executeQueue(memory, queries)
}

func executeQueue(memory []int, queries [][]int) []int {

}
