package main

/**
@author Shitanford
@create 2021-10-01 15:36
*/

func main() {

}

func flatten(root *TreeNode)  {
	for root != nil {
		if root.Left != nil {
			getRightLeaf(root.Left).Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}

func getRightLeaf(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}