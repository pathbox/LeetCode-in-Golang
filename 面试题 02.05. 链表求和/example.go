package LeetCode0205

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	dummy := &ListNode{}
	result := dummy
	carry := 0 // 进位
	for l1 != nil || l2 != nil {
		resultV := carry
		if l1 != nil {
			resultV += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			resultV += l2.Val
			l2 = l2.Next
		}
		result.Next = &ListNode{Val: resultV % 10}
		result = result.Next
		carry = resultV / 10
	}
	if carry != 0 {
		result.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
