package LeetCode138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// use two map 第一个map为了copy 新的Node 第二个map为了将新的Node 的Next Random连接起来
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	m := make(map[*Node]*Node) // map的key是原指针的节点，val是copy的temp新节点
	cur := head
	for cur != nil {
		temp := &Node{Val: cur.Val}
		m[cur] = temp
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}
