package LeetCode021

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head = &ListNode{Val: -1}
	var pre = head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	} else if l2 != nil {
		pre.Next = l2
	}
	return head.Next
}
