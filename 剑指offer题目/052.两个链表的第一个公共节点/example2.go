package Offer052

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, d := headA, headB
	for p != d {
		if p == nil && d == nil {
			return p
		}

		if p != nil {
			p = p.Next
		} else {
			p = headB // 指针走到另一个链表
		}

		if d != nil {
			d = d.Next
		} else {
			d = headA // 指针走到另一个链表
		}
	}

	return p
}
