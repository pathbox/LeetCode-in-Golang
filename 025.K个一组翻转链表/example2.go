package LeetCode025

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}

	cur := head
	var start, pre, next *ListNode
	count := 1
	for cur != nil {
		next = cur.Next
		if count == k {
			if pre == nil {
				start = head
			} else {
				start = pre.Next
			}
			if pre == nil {
				head = cur
			}
			reverse(pre, start, cur, next)
			pre = start
			count = 0
		}
		count++
		cur = next
	}
	return head
}

func reverse(left, start, end, right *ListNode) {
	pre := start
	cur := start.Next
	for cur != right {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	if left != nil {
		left.Next = end
	}
	start.Next = right
}
