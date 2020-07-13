package LeetCode148

type ListNode struct {
	Val  int
	Next *ListNode
}

// 单链表快速排序Time Complexity: O(nlogn), Space: O(n)
func sortList(head *ListNode) *ListNode {
	// 如果 head为空或者head就一位,直接返回
	if head == nil || head.Next == nil {
		return nil
	}
	quickSort(head, nil)
	return head
}
