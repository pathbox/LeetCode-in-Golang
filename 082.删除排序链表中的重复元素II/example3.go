package LeetCode082

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}

	a := dummy
	b := head

	for b != nil && b.Next != nil {
		if a.Next.Val != b.Next.Val {
			a = a.Next
			b = b.Next
		} else {
			//如果a、b指向的节点值相等，就不断移动b，直到a、b指向的值不相等
			for b != nil && b.Next != nil && a.Next.Val == b.Next.Val {
				b = b.Next
			}
			a.Next = b.Next
			b = b.Next
		}
	}
	return dummy.Next
}
