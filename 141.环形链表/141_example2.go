package LeetCode021

// 快慢指针为啥可以呢？因为如果是一个环一个人在前面走的更快，另一个慢，最终快的都会赶上慢的，比慢的多一圈，如果不是环则快的最先到头，选择2个步长，选择三个有可能会刚好跳过
func hasCycle(head *ListNode) bool { // 快慢指针。假如爱有天意，那么快慢指针终会相遇
	if head == nil {
		return false
	}

	fastHead := head.Next // 快指针，每次走两步
	for fastHead != nil && head != nil && fastHead.Next != nil {
		if fastHead == head { // 快慢指针相遇，表示有环
			return true
		}
		fastHead = fastHead.Next.Next
		head = head.Next
	}
	return false
}
