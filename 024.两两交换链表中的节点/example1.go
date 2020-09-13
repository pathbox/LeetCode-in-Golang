package LeetCode024

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/solution/24_swap-nodes-in-pairs-by-hawflakes-w/

// 递归解法
func swapPairs(head *ListNode) *ListNode {
	// 第一个节点和第二个节点为空时，直接返回原链表
	if head == nil || head.Next == nil {
		return head
	}

	// 声明firstNode，secondNode，并将head赋给firstNode，head的下一个节点赋给secondNode
	firstNode := head
	secondNode := head.Next

	// 交换firstNode与second,
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}
