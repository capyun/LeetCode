package main

import "fmt"

/**
@author Shitanford
@create 2021-10-01 20:08
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Left: &TreeNode{

		},
	}
	fmt.Println(diameterOfBinaryTree(root))
}

var maxLen int

func diameterOfBinaryTree(root *TreeNode) int {
	maxLen = 0
	depth(root)
	return maxLen
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)
	maxLen = maxInt(maxLen, leftDepth + rightDepth)
	return maxInt(leftDepth, rightDepth) + 1
}

func maxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}