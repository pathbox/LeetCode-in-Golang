package LeetCode061

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, n := head, 1
	for cur.Next != nil {
		cur = cur.Next
		n++ // 得到链表长度
	}
	cur.Next = head // 形成环
	m := n - k%n    // 找到断开的点位置
	for i := 0; i < m; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil // 把环断开
	return head
}
