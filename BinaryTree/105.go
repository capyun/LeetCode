package main

import "fmt"

/**
@author Shitanford
@create 2021-09-30 21:09
*/

func main() {
	pre := []int{3,9,20,15,7}
	in := []int{9,3,15,20,7}
	root := buildTree(pre, in)
	fmt.Println(root)
	printTree(root)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode)  {
	arr := make([]*TreeNode, 8)
	arr[1] = root
	for i := 1; i < 3; i++ {
		arr[i*2] = arr[i].Left
		arr[i*2+1] = arr[i].Right
	}
	for _, v := range arr {
		if v != nil {
			fmt.Printf("%d ", v.Val)
		} else {
			fmt.Printf("nil ")
		}
	}
	fmt.Println()
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(inorder) == 0 {
//		return nil
//	}
//	root := &TreeNode{Val: preorder[0]}
//	index := find(inorder, preorder[0])
//	root.Left = buildTree(preorder[1:], inorder[: index])
//	root.Right = buildTree(preorder[index+1: ], inorder[index+1: ])
//	return root
//}
//
//func find(arr []int, value int) int {
//	for i, v := range arr {
//		if v == value {
//			return i
//		}
//	}
//	return -1
//}

func buildTree(preorder []int, inorder []int) *TreeNode {
	index := make(map[int]int)
	for i, v := range inorder {
		index[v] = i
	}
	return arrayToTree(preorder, index, 0)
}

func arrayToTree(preorder []int, index map[int]int, rootIndex int) *TreeNode {
	i := index[preorder[rootIndex]]
	if rootIndex == i {
		return nil
	}
	root := &TreeNode{Val: preorder[rootIndex]}
	root.Left = arrayToTree(preorder, index, rootIndex+1)
	root.Right = arrayToTree(preorder, index, i+1)
	return root
}