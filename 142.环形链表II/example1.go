package LeetCode142

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode/
func detectCycle(head *ListNode) *ListNode {
	// 如果直接nil或者下一个节点为nil,那么就肯定没有环了
	if head == nil || head.Next == nil {
		return nil
	}

	// 定义两个节点,一个为快指针,一个慢指针
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		// 快指针肯定跑得快,快指针如果走到了nil,那么说明没有环
		if fast.Next == nil {
			return nil
		}
		// 快指针每次跳两步
		fast = fast.Next.Next
		// 慢指针每次跳一步
		slow = slow.Next
		//  当 快指针和慢指针相遇了
		if fast == slow {
			// head 节点开始走,fast指针改成每次走一步,相遇的时候 就是环形的起点
			// 相遇的地方到初始节点处的距离就是起点到环的入口节点的距离
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}
	}
	return nil
}
