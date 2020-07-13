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

// 用于交换两个节点上的数值
func swap(a, b *ListNode) {
	a.Val, b.Val = b.Val, a.Val
}

// 快速排序的主体函数
func quickSort(head, end *ListNode) {
	if head == end || head.Next == end {
		return // 头节点已经是结束节点或要排序只有一个头节点直接返回
	}
	pivot := head.Val             // 头节点的值作为pivot
	slow, fast := head, head.Next // 定义快慢指针
	for fast != end {             // 当快指针不等于结束指针时，不断执行以下操作
		if fast.Val <= pivot { // 如果快指针小于pivot
			slow = slow.Next // 不断向后移动慢指针
			swap(slow, fast) // 交换快慢指针指向的节点值
		}
		fast = fast.Next // 快指针向后移动一位
	}
	// 循环结束后交换头节点和慢指针指向的值,把pivot放在大小两堆数值的中间
	swap(head, slow)
	quickSort(head, slow) // 递归处理pivot左右两边的链表
	quickSort(slow.Next, end)
}
