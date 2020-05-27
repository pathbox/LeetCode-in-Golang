package LeetCode445

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []int{}, []int{}

	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	prev := &ListNode{}
	carry := 0

	len1 := len(stack1)
	len2 := len(stack2)

	for len1 != 0 || len2 != 0 || carry != 0 { // carry != 0 还有最后一个进位的可能情况，此时两个stack已经遍历完了
		temp1, temp2 := 0, 0
		if len1 != 0 {
			len1--
			temp1 = stack1[len1]
		}
		if len2 != 0 {
			len2--
			temp2 = stack2[len2]
		}
		prev.Next = &ListNode{(temp1 + temp2 + carry) % 10, prev.Next} // prev 节点保持不变，每次新的节点插入到prev和prev.Next之间，变为prev.Next 插入到已存在节点的前面
		carry = (temp1 + temp2 + carry) / 10
	}
	return prev.Next
}
