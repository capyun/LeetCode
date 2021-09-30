package main

/**
@author Shitanford
@create 2021-09-27 15:12
*/

func main() {

}

func sumNumbers(root *TreeNode) int {
	return sumNumbers2(root, 0)
}

func sumNumbers2(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}
	x = x*10 + root.Val
	if root.Left != nil || root.Right != nil {
		return sumNumbers2(root.Left, x) + sumNumbers2(root.Right, x)
	} else {
		return x
	}
}