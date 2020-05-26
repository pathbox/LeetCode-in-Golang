package LeetCode002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	node1 := l1
	node2 := l2
	l3 := new(ListNode)
	node3 := l3
	a := 0
	for node1 != nil || node2 != nil || a > 0 {
		node3.Next = new(ListNode)
		node3 = node3.Next
		b := 0
		c := 0
		if node1 != nil {
			b = node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			c = node2.Val
			node2 = node2.Next
		}
		node3.Val = (a + b + c) % 10 // 去除进位的值
		a = (a + b + c) / 10         // 进位值
	}
	return l3.Next
}
