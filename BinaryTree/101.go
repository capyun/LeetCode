package main

import "fmt"

/**
@author Shitanford
@create 2021-10-01 15:57
*/

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}
	isSymmetric(root)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelp(root.Left, root.Right)
}

func isSymmetricHelp(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricHelp(left.Left, right.Right) &&
		isSymmetricHelp(left.Right, right.Left)
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{101}
	}
	arr := inOrder(root.Left)
	arr = append(arr, root.Val)
	arr = append(arr, inOrder(root.Right)...)
	return arr
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}