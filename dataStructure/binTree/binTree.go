package binTree

/**
@author Shitanford
@create 2021-04-15 11:08
*/

import (
	"DS.com/stack"
	"fmt"
)

type BinTree struct {
	value int
	lChild, rChild *BinTree
	nodeNum int
}

func NewBT(value int) *BinTree {
	return &BinTree{value: value,
					lChild: nil,
					rChild: nil,
					nodeNum: 1}
}

func (bt *BinTree) GetValue() int {
	return bt.value
}

func (bt *BinTree) GetLeft() *BinTree {
	return bt.lChild
}

func (bt *BinTree) GetRight() *BinTree {
	return bt.rChild
}

func (bt *BinTree) SetLeft(tree *BinTree)  {
	bt.lChild = tree
}

func (bt *BinTree) SetRight(tree *BinTree)  {
	bt.rChild = tree
}

func (bt *BinTree) PreTraverse() {
	n := bt
	s := stack.New(bt.nodeNum)
	for {
		for n != nil {
			fmt.Printf("%v", n.value)
			if n.rChild != nil {
				s.Push(n.rChild)
			}
			n = n.lChild
		}
		if s.Empty() {
			break
		}
		n, _ = (s.Pop()).(*BinTree)
	}
}

func (bt *BinTree) InTraverse()  {
	n := bt
	s := stack.New(bt.nodeNum)
	for {
		for n != nil {
			s.Push(n)
			n = n.lChild
		}
		if s.Empty() {
			break
		}
		n, _ = (s.Pop()).(*BinTree)
		fmt.Printf("%v", n.value)
		n = n.rChild
	}
}

func (bt *BinTree) PostTraverse()  {
	n := bt
	s := stack.New(bt.nodeNum)
	for {
		for n != nil {
			s.Push(n)
			if n.lChild != nil {
				n = n.lChild
			} else {
				n = n.rChild
			}
		}
		n, _ = (s.Pop()).(*BinTree)
		fmt.Printf("%v", n.value)
		if s.Empty() {
			break
		}
		if (s.Top()).(*BinTree).lChild == n {
			n = (s.Top()).(*BinTree).rChild
		} else {
			//右节点访问过了，则继续出栈，（给个空值）
			n = nil
		}
	}
}

func (bt *BinTree) search(value int) *BinTree {
	ptr := bt
	for ptr != nil {
		if value < ptr.GetValue() {
			ptr = ptr.GetLeft()
		} else if value > ptr.GetValue() {
			ptr = ptr.GetRight()
		} else {
			return ptr
		}
	}
	return nil
}

type BST struct {
	BinTree
}

func NewBST(value int) *BST {
	return &BST{*NewBT(value)}
}

func (bst *BST) Search(value int) *BST {
	bt := &bst.BinTree
	ans := bt.search(value)
	return BST(ans)
}

func (bst *BST) Insert(value int) *BST {
	if bst.BinTree == nil {
		return NewBST(value)
	}
	ptr := bst
	fmt.Printf("%v\n", ptr.BinTree)
	if value < ptr.GetValue() {
		ptr.BinTree = ptr.GetLeft()
		bst.SetLeft(ptr.Insert(value).BinTree)
	} else if value > ptr.GetValue() {
		ptr.BinTree = ptr.GetRight()
		bst.SetRight(ptr.Insert(value).BinTree)
	}
	return bst
}

//右儿子为空，可以不管，
//不管右儿子为不为空，都要访问，这是个循环结构的问题