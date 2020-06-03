package Offer006

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return appendData(head)
}

func appendData(head *ListNode) []int {
	if head.Next != nil {
		list := appendData(head.Next) // 递归，不断的遍历到最后一个节点，之后出栈，先进入list的就是最后一个节点值，递归不断出栈，相当于倒序了
		list = append(list, head.Val)
		return list
	}
	return []int{head.Val}
}
