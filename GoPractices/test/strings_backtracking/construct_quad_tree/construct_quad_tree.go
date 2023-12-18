package main

import "fmt"

/*
https://leetcode.com/problems/construct-quad-tree/?envType=study-plan-v2&envId=top-interview-150
Given a n * n matrix grid of 0's and 1's only. We want to represent grid with a Quad-Tree.

Return the root of the Quad-Tree representing grid.

A Quad-Tree is a tree data structure in which each internal node has exactly four children. Besides, each node has two attributes:

    val: True if the node represents a grid of 1's or False if the node represents a grid of 0's. Notice that you can assign the val to True or False when isLeaf is False, and both are accepted in the answer.
    isLeaf: True if the node is a leaf node on the tree or False if the node has four children.

class Node {
    public boolean val;
    public boolean isLeaf;
    public Node topLeft;
    public Node topRight;
    public Node bottomLeft;
    public Node bottomRight;
}

We can construct a Quad-Tree from a two-dimensional area using the following steps:

    If the current grid has the same value (i.e all 1's or all 0's) set isLeaf True and set val to the value of the grid and set the four children to Null and stop.
    If the current grid has different values, set isLeaf to False and set val to any value and divide the current grid into four sub-grids as shown in the photo.
    Recurse for each of the children with the proper sub-grid.



*/

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)
	return helper(0, 0, n, n, grid)
}

func helper(rowS, colS, rowE, colE int, grid [][]int) *Node {
	if rowS >= rowE || colS >= colE {
		return nil
	}
	root := &Node{}
	isLeaf := true
	for i := rowS; i < rowE; i++ {
		for j := colS; j < colE; j++ {
			if grid[rowS][colS] != grid[i][j] {
				isLeaf = false
				break
			}
		}
		if !isLeaf {
			break
		}
	}
	if isLeaf {
		root.Val = (grid[rowS][colS] == 1)
		root.IsLeaf = isLeaf
		return root
	}
	root.Val = false

	rowMid := (rowS + rowE) / 2
	colMid := (colS + colE) / 2
	root.TopLeft = helper(rowS, colS, rowMid, colMid, grid)
	root.TopRight = helper(rowS, colMid, rowMid, colE, grid)
	root.BottomLeft = helper(rowMid, colS, rowE, colMid, grid)
	root.BottomRight = helper(rowMid, colMid, rowE, colE, grid)
	return root
}

func main() {
	grid := [][]int{{1, 1, 0, 0}, {0, 0, 1, 1}, {1, 1, 0, 0}, {0, 0, 1, 1}}
	fmt.Println(construct(grid))
	fmt.Println("end")
}
