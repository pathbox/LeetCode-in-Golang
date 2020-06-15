package Offer035

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 创建新的节点， 修改节点的关系， 生成节点 要分开做。

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head

	//创建新的节点，在原来节点的后面
	for cur != nil {
		new_node := new(Node)
		new_node.Val = cur.Val
		new_node.Next = cur.Next

		cur.Next = new_node
		cur = new_node.Next
	}

	//改变新节点的Random 指向关系
	cur = head
	for cur != nil {
		if cur.Random == nil {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	newHead := new(Node)

	p := newHead
	cur = head

	for cur != nil {
		p.Next = cur.Next
		p = p.Next

		cur.Next = p.Next
		cur = cur.Next
	}
	return newHead.Next
}
