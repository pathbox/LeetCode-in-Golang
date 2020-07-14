

type ListNode struct {
	Val  int
	Next *ListNode
}

// 单链表归并排序 Time: O(n*logn) Space: O(logn)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head // 快慢指针初始化在链表头
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 循环结束后慢指针指向的是前一半链表的最后一个元素
	right := sortList(slow.Next)            // 先排序后一段链表
	slow.Next = nil                         // 然后把慢指针的Next置为空指针
	left := sortList(head)                  // 再排序前一段链表
	return mergeTwoSortedLists(left, right) // 合并排序后的两段链表
}

func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy // dummy指针没有移动，一直移动的是p
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}