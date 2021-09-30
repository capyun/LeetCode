package main

import "fmt"

/**
@author Shitanford
@create 2021-05-11 16:21
*/

type queue struct {
	arr[] int
	head, tail int
}

func newQueue(n int) queue {
	return queue{arr: make([]int, n+1),
				head:0, tail:0}
}
func (q *queue) push(v int)  {
	q.arr[q.tail] = v
	q.tail+=1
}
func (q *queue) pop() int {
	q.head++
	return q.arr[q.head-1]
}
func (q queue) isEmpty() bool {
	return q.head == q.tail
}

type node struct {
	next *node
	num int
}
func main() {
	dependencies := [][]int{{2,1},{3,1},{1,4}}
	n, k := 4, 2
	fmt.Printf("%d", minNumberOfSemesters(n, dependencies, k))
}

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	ans := 0
	q := newQueue(n)
	arr := make([]node, n+1)
	for i:=0; i<len(dependencies); i++ {
		newNode := new(node)
		newNode.num = dependencies[i][1]
		newNode.next = arr[dependencies[i][0]].next
		arr[dependencies[i][0]].next = newNode
		arr[dependencies[i][1]].num += 1
	}
	var ptr *node
	learned := 0
	for i := 1; i <= n; i++ {
		if arr[i].num == 0 {
			q.push(i)
			arr[i].num = -1
		}
	}
	for learned < n {
		for i := 0; i < k && !q.isEmpty(); i++ {
			course := q.pop()
			learned ++
			ptr = arr[course].next
			for ptr != nil {
				arr[ptr.num].num--
				ptr = ptr.next
			}
		}
		ans ++
		for i := 1; i <= n; i++ {
			if arr[i].num == 0 {
				q.push(i)
				arr[i].num = -1
			}
		}
	}
	return ans
}