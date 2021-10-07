package main

/**
@author Shitanford
@create 2021-10-03 16:12
*/

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

func helper(root *TreeNode)  {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	helper(root.Left)
	helper(root.Right)
}