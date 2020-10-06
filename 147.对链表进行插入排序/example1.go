package LeetCode147

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead, pre, node := &ListNode{}, head, head.Next
	newHead.Next = head
	for node != nil {
		if node.Val < pre.Val { // 后面比前面的值小
			temp := newHead
			for temp.Next.Val < node.Val { // 将node插入到newHead中
				temp = temp.Next
			}
			// 将pre放到newHead中应该放的位置，同时node指向下一个节点
			pre.Next = node.Next  // 删除node节点原来位置
			node.Next = temp.Next // 将node插入到temp后面
			temp.Next = node
			node = pre.Next // 继续往下遍历还没有遍历的node节点
		} else { // 已经有序
			pre, node = pre.Next, node.Next
		}
	}
	return newHead.Next
}
