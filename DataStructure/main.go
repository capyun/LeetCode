package main

import (
	"fmt"
	"shitanford.com/leetcode/datastruct/heap"
)

/**
@author Shitanford
@create 2021-05-14 11:24
*/

func main() {
	arr := []int{5,1,8,10,7}
	fmt.Printf("%d", lastStoneWeight(arr))
}

func lastStoneWeight(stones []int) int {
	h := heap.NewHeap(stones)
	for h.GetSize()>1 {
		x := h.Pop()
		y := h.Pop()
		if x > y {
			h.Push(x - y)
		}
	}
	if h.GetSize()==1 {
		return h.Pop()
	} else {
		return 0
	}
}