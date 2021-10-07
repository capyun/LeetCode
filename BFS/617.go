package BFS

/**
@author Shitanford
@create 2021-10-03 16:17
*/

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

}

func helper(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = helper(root1.Left, root2.Left)
	root1.Right = helper(root1.Right, root2.Right)
	return root1
}