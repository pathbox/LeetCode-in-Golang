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

	current := dummy
	var delvar int //记录删除的值
	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			delvar = current.Next.Val //记录删除的元素
			//删除相同的元素
			for current.Next != nil && current.Next.Val == delvar {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}

	}
	return dummy.Next
}
