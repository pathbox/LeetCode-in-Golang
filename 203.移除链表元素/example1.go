package LeetCode203

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre, cur := dummy, head
	for cur != nil {
		if cur.Val == val { // cur就是要删除的节点
			pre.Next = cur.Next
		} else {
			pre = cur // pre 往前移动
		}
		cur = cur.Next // cur往前移动
	}
	return dummy.Next
}
