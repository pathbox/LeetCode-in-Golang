package Offer006

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var newHead *ListNode // 新增一个新的头节点
	res := []int{}
	for head != nil { // 不断遍历head指针
		node := head.Next   // 保存next节点
		head.Next = newHead // 倒序连接
		newHead = head      // 将newHead移到头部 并不是说链表断开了，指针之间就不能赋值，指针是否能赋值和移动，和链表的连接性无关
		head = node
	}

	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}
	return res
}
