package Offer018

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head} // 假头
	first := dummyHead                         // 双指针1
	second := dummyHead.Next                   // 双指针2
	for second != nil {
		if second.Val == val { // 删除的是second节点
			first.Next = second.Next
			break
		}
		first = first.Next
		second = second.Next
	}
	return dummyHead.Next
}
