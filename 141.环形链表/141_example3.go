package LeetCode021

//  O(1) 的空间复杂度
func hasCycle(head *ListNode) bool { // 走自己的路让别人无路可走思路
	for head != nil {
		if head.Val == 18464187 { // 这是自己走过的路，说明有环 如果Next的值是预先定义标记的值，说明有环
			return true
		}
		head.Val = 18464187
		head = head.Next
	}
	return false
}
