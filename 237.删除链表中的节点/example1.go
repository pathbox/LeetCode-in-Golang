package LeetCode237

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next // node节点被Next覆盖，删除重复的一个Next节点
}

/*
这道题没有给出链表的头节点，而是直接给出要删除的节点，让我们进行原地删除
实际上是将Next节点覆盖为要删除的node节点，然后把重复的Next节点删除一个，来实现原地删除node节点的效果
*/
