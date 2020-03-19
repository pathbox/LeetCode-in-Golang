package LeetCode021

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 如果有一条为nil， 直接返回另一条
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 定义头head和移动指针node
	var head, node *ListNode
	// 选小的那条开始
	// 前面判断过l1和l2了，所以这里不可能是空指针
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}
	// 循环比较，同时更新
	// 始终选择小的值连到node上
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			// 指向小的那条
			node.Next = l1
			// 小的往前走
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		// 【重要】node往前走
		node = node.Next
	}
	// 循环完了，把剩下的直接放到后面，因为本身就是有序的，所以接在后面就行
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	// 返回最初的那个头
	// 这个头不能动，否则找不到了
	return head
}
