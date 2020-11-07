package offer062

// 约瑟夫环

type Node struct {
	Next *Node
}

func josephusKill(head *Node, m int) *Node {
	if head == nil || head.Next == head || m < 1 {
		return head
	}

	cur := head.Next
	tmp := 1
	for cur != head {
		tmp++
		cur = cur.Next
	}
	tmp = getLive(tmp, m)
	for tmp != 0 {
		head = head.Next
		tmp--
	}
	head.Next = head
	return head
}

func getLive(i, m int) int {
	if i == 1 {
		return 1
	}
	return (getLive(i-1, m)+m-1)%i + 1
}
