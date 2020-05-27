package LeetCode445

type ListNode struct {
	Val  int
	Next *ListNode
}

// 翻转链表法
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := reverse(l1)
	node2 := reverse(l2)

	prev := &ListNode{}
	a := 0
	for node1 != nil || node2 != nil || a > 0 {
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
		prev.Next = &ListNode{(b + c + a) % 10, prev.Next}
		a = (a + b + c) / 10 // 进位值
	}

	return prev.Next
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	for node.Next != nil {
		prev, node.Next, node = node, prev, node.Next

		if node.Next == nil {
			node.Next = prev
			break
		}
	}
	return node
}
