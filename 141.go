package main

/**
@author Shitanford
@create 2021-09-28 10:49
*/

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//func hasCycle(head *ListNode) bool {
//	var count = 0
//	for head != nil && count <= 10000 {
//		head = head.Next
//		count++
//	}
//	return count > 10000
//}

func hasCycle(head *ListNode) bool {
	nodeSet := make(map[*ListNode]int, 0)
	for head != nil {
		if _, exists := nodeSet[head]; exists {
			return true
		}
		nodeSet[head] = 0
		head = head.Next
	}
	return false
}