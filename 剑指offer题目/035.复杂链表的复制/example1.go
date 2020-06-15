package Offer035

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var visit = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
	return dfs(head)
}

func dfs(head *Node) *Node {
	if head == nil {
		return nil
	}
	if newNode, ok := visit[head]; ok {
		return newNode
	}
	copyNode := &Node{
		Val: head.Val,
	}
	visit[head] = copyNode
	copyNode.Next = dfs(head.Next)
	copyNode.Random = dfs(head.Random)
	return copyNode
}
