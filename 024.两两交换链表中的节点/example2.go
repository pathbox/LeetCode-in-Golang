package LeetCode024

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy

	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		prev.Next = next

		prev = head
		head = head.Next
	}
	return dummy.Next
}
