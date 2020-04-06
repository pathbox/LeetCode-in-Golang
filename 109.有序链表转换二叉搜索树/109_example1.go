package LeetCode109

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid := findMiddle(head) // find root
	node := &TreeNode{      // mid的值就是当前节点
		Val: mid.Val,
	}

	if head == mid { // just one element
		return node
	}

	node.Left = sortedListToBST(head)      // 左部分继续查找链表
	node.Right = sortedListToBST(mid.Next) // 右部分继续查找链表
	return node
}

func findMiddle(node *ListNode) *ListNode {
	var prev *ListNode
	slow, fast := node, node
	for fast != nil && fast.Next != nil { // 链表中，怎么找到中间位置节点，利用快慢指针，快指针是慢指针的两倍速度，当快指针为nil了，此时slow指针就是走到中间节点位置
		prev = slow
		slow = slow.Next      // move 1
		fast = fast.Next.Next // move 2
	}

	if prev != nil {
		prev.Next = nil
	}
	return slow
}
