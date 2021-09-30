package main

import "fmt"

/**
@author Shitanford
@create 2021-04-06 22:07
*/

func main() {
	//[2,4,3]
	//[5,6,4]
	l1 := new(ListNode)
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4

	l2 := new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6

	ans := addTwoNumbers(l1, l2)
	for ans != nil {
		fmt.Printf("%d ", ans.Val)
		ans = ans.Next
	}
}

type ListNode struct {
	 Val int
	 Next *ListNode
 }

//func generateList(list *ListNode) {
//}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	ans := new(ListNode)
//	ptr := ans
//	carry := 0
//	for l1 != nil && l2 != nil {
//		ptr.Next = new(ListNode)
//		ptr = ptr.Next
//		ptr.Val = l1.Val + l2.Val + carry
//		l1 = l1.Next
//		l2 = l2.Next
//		carry = 0
//		if ptr.Val > 9 {
//			ptr.Val -= 10
//			carry = 1
//		}
//	}
//	if l1 == nil {
//		l1 = l2
//	}
//	for l1 != nil {
//		ptr.Next = new(ListNode)
//		ptr = ptr.Next
//		ptr.Val = l1.Val + carry
//		l1 = l1.Next
//		carry = 0
//		if ptr.Val > 9 {
//			ptr.Val %= 10
//			carry = 1
//		}
//	}
//	if carry == 1 {
//		ptr.Next = new(ListNode)
//		ptr = ptr.Next
//		ptr.Val = 1
//	}
//	return ans.Next
//}

