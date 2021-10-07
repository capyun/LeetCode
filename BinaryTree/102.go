package main

/**
@author Shitanford
@create 2021-10-01 20:34
*/

func main() {

}

//func levelOrder(root *TreeNode) [][]int {
//	var ans [][]int
//	var nodeSet []*TreeNode
//	if root == nil {
//		return ans
//	}
//	nodeSet = append(nodeSet, root)
//
//	i, j := 0, 1
//	for i < j {
//		var arr []int
//		for k := j; i < k; i++ {
//			arr = append(arr, nodeSet[i].Val)
//
//			if nodeSet[i].Left != nil {
//				nodeSet = append(nodeSet, nodeSet[i].Left)
//				j++
//			}
//			if nodeSet[i].Right != nil {
//				nodeSet = append(nodeSet, nodeSet[i].Right)
//				j++
//			}
//		}
//		ans = append(ans, arr)
//	}
//	return ans
//}

func levelOrder(root *TreeNode) [][]int {
	var arr [][]int
	levelOrderHelper(root, &arr, 0)
	return arr
}

func levelOrderHelper(root *TreeNode, arr *[][]int, level int)  {
	if root == nil {
		return
	}
	if len(*arr) == level {
		*arr = append(*arr, []int{})
	}
	(*arr)[level] = append((*arr)[level], root.Val)
	levelOrderHelper(root.Left, arr, level+1)
	levelOrderHelper(root.Right, arr, level+1)
}
