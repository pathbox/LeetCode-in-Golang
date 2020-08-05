package LeetCode026

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur, next := head, head.Next

	for next != nil {
		if cur.Val == next.Val {
			next = next.Next
		} else {
			cur.Next = next // 将所有与cur节点重复的节点跳过
			cur = next      // 继续移动两指针
			next = next.Next
		}
	}
	cur.Next = nil
	return head
}

//2.0 下面是删除 ***无序链表*** 中重复的元素 (刚开始题目没看仔细，就一起发出来了)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var curr, next = head, head.Next

	for next != nil {
		if next.Val != head.Val {
			next = next.Next
			curr = curr.Next
		} else {
			next = next.Next
			curr.Next = next
		}
	}
	if head.Next != nil {
		deleteDuplicates(head.Next) // 每个节点都和后面的所有节点比较
	}
	return head
}
