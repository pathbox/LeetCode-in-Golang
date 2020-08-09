package LeetCode143

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
首先重新排列后，链表的中心节点会变为最后一个节点。所以需要先找到链表的中心节点：876. 链表的中间结点
可以按照中心节点将原始链表划分为左右两个链表。
2.1. 按照中心节点将原始链表划分为左右两个链表，左链表：1->2->3 右链表：4->5。
2.2. 将右链表反转，就正好是重排链表交换的顺序，右链表反转：5->4。反转链表：206. 反转链表
合并两个链表，将右链表插入到左链表，即可重新排列成：1->5->2->4->3.
*/

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	left := head
	right := slow.Next // 中间节点
	slow.Next = nil
	right = Reverse(right) // 反转右部分链表

	for right != nil {
		lNext := left.Next
		rNext := right.Next
		left.Next = right
		right.Next = lNext
		left = lNext
		right = rNext
	}
}

func Reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{}
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
