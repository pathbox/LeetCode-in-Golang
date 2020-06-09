package Offer018

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	// 情况1
	if head.Val == val {
		return head.Next
	}

	// 情况2
	originHead := head
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}

	return originHead
}
