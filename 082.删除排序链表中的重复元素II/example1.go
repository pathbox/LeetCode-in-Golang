package LeetCode082

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
遍历所有节点，比对一下当前值和之前看见的值一不一样，如果不一样，并且当前值和下一个节点的值也不一样，那么这个节点就值得加入到返回结果的链表中。
唯一主要注意的两点是，最后一个节点没法和它自己的Next来比较Val，因为它的Next是nil, 因此如果最后一个节点的Val之前没见过，那么这个节点也值得加入到返回结果链表中。

最后记得把返回结果链表当前位置的Next设置成nil
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ret := &ListNode{
		Next: head,
	}

	c := ret
	seen := head.Val - 1
	for l := head; l != nil; l = l.Next {
		switch {
		case l.Val != seen && l.Next != nil && l.Val != l.Next.Val:
			fallthrough
		case l.Next == nil && l.Val != seen:
			c.Next = l
			c = c.Next
		}
		seen = l.Val
	}
	c.Next = nil
	return ret.Next
}
