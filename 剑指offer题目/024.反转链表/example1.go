package Offer024

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre, cur := head, &ListNode{}
	for pre != nil {
		tmpNext := pre.Next
		pre.Next = cur
		cur = pre
		pre = tmpNext
	}
	return cur
}
