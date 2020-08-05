package LeetCode086

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	greatDummy := &ListNode{}
	lessDummy := &ListNode{}
	curGreat := greatDummy
	curLess := lessDummy
	for head != nil {
		if head.Val >= x {
			curGreat.Next = head
			curGreat = head
		} else {
			curLess.Next = head
			curLess = head
		}
		head = head.Next
	}
	curGreat.Next = nil
	curLess.Next = greatDummy.Next
	return lessDummy.Next
}
