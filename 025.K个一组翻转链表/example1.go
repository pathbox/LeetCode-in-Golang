package LeetCode025

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ { // 每k个一组
			tail = tail.Next // tail 往前走k步
			if tail == nil { // 到结尾了
				return dummy.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail) // head -> tail 一段给反转了
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return dummy.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next // prev一开始设为tail.Next，想想prev设为nil的情况
	cur := head
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return tail, head // 这一段反转后返回
}
