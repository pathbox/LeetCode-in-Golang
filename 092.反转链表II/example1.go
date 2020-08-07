package LeetCode092

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummy := &ListNode{Next: head} // 这样初始化dummy
	prem := dummy
	for i := 1; i <= m-1; i++ {
		prem = prem.Next // 移到m的前一个
	}
	cur, next := prem.Next, prem.Next.Next // cur就是第m个节点
	var pre, nextnext *ListNode
	// 需要三个指针进行反转
	for i := m; i < n; i++ {
		nextnext = next.Next
		cur.Next, pre, next.Next = pre, cur, cur
		cur, next = next, nextnext
	}
	prem.Next.Next, prem.Next = next, cur
	return dummy.Next
}
