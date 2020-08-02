package LeetCode019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	//建立哨兵节点
	dummy := &ListNode{}
	dummy.Next = head
	preN := dummy
	i := 1
	for head != nil {
		if i > n { // 在n次遍历前，head先移动n次
			preN = preN.Next
		}
		head = head.Next
		i++
	}
	//跨过倒数第N个节点 preN 是要删除节点的前一个
	preN.Next = preN.Next.Next
	return dummy.Next
}
