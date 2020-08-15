package LeetCode234

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先用快慢指针找到中间点， 然后把后半段链表反转， 再和原链表从头比较值
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next == nil { //找中点
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil { //奇数链表取下一点
		slow = slow.Next
	}

	var pre *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = pre
		pre = slow
		slow = tmp
	}

	for pre != nil { //和原来的比较 此时的pre就是最后一个节点 即反转后链表的头结点
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}

//总共O(2n） 空间O(2)
