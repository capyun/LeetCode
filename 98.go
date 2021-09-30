package main

import "math"

/**
@author Shitanford
@create 2021-09-27 10:21
*/

func main() {
	root := &TreeNode{Val: 0}
	//root.Left = &TreeNode{Val: 1}
	//root.Right = &TreeNode{Val: 4}
	isValidBST(root)
}

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

//func generateTree() *TreeNode {
//	for fmt.Scanf()
//}

var prevVal int = math.MinInt32

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST(root.Left) {
		return false
	}
	if root.Val <= prevVal {
		return false
	}
	prevVal = root.Val
	if !isValidBST(root.Right) {
		return false
	}
	return true
}

//func isValidBST(root *TreeNode) bool {
//	return isValidBST2(root, math.MinInt32, math.MaxInt32)
//}

//func isValidBST2(root *TreeNode, min int, max int) bool {
//	if root == nil {
//		return true
//	}
//	if root.Val <= min || root.Val >= max {
//		return false
//	}
//	return isValidBST2(root.Left, min, root.Val) &&
//		isValidBST2(root.Right, root.Val, max)
//}