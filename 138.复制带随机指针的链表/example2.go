package LeetCode138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val: cur.Val}
		cur.Next.Next = next
		cur = next
	}

	cur = head
	for cur != nil {
		next := cur.Next.Next
		curCopy := cur.Next
		if cur.Random != nil {
			curCopy.Random = cur.Random.Next
		}
		cur = next
	}
	res := head.Next
	cur = head
	for cur != nil {
		next := cur.Next.Next
		curCopy := cur.Next
		cur.Next = next
		if next != nil {
			curCopy.Next = next.Next
		}
		cur = next
	}
	return res
}
