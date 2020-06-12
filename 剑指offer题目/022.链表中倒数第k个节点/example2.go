package Offer022

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var res []*ListNode
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	l := len(res)
	if l >= k {
		return res[l-k]
	}
	return nil
}
