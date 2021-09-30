package main

/**
@author Shitanford
@create 2021-09-26 15:58
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	ans := new(ListNode)
	i := len(stack1)-1
	j := len(stack2)-1
	sum := 0
	carry := 0
	for i >= 0 || j >= 0 || carry > 0 {
		sum = carry
		if i >= 0 {
			sum += stack1[i]
			i--
		}
		if j >= 0 {
			sum += stack2[j]
			j--
		}
		if sum > 9 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		ptr := &ListNode{Val: sum, Next: ans.Next}
		ans.Next = ptr
	}
	return ans.Next
}
