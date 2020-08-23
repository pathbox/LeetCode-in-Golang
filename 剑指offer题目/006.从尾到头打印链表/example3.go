package Offer006

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将链表数据放入数组，再倒序输出
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
