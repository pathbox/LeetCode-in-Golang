package Offer022

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针方式：间隔k，当fast指针到达末尾，slow指针就是倒数第k个
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for k > 0 && fast != nil { // 先让fast走k部
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
