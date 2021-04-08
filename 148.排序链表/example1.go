package LeetCode148

type ListNode struct {
	Val  int
	Next *ListNode
}

// 归并排序
func sortList(head *ListNode) *ListNode {
	// 为空或者只有一个元素
	if head == nil || head.Next == nil {
		return head
	}
	// 1. 找中点，二分，左右分别进行排序,快慢指针 fast比slow快一个是为了让slow.next可以成为中点，便于截断
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 注意要将链表截断
	rNext := slow.Next
	slow.Next = nil
	l := sortList(head)
	r := sortList(rNext)
	// 2. 合并，将两个list合并为一个
	return mergeList(l, r)
}

// 因为要常数级的空间复杂度，所以这里把 temphead传进来
func mergeList(l, r *ListNode) *ListNode {
	temphead := &ListNode{Val: 0}
	p := temphead
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l != nil {
		p.Next = l
	}
	if r != nil {
		p.Next = r
	}
	return temphead.Next // 返回的是temphead.Next 而不是temphead
}
