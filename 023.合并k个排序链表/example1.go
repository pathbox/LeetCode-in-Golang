package LeetCode023

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
解题思路 - 分治算法
像归并排序一样使用分治法处理，假如k个链表分别为a1, a2, a3,…, ak

不断递归直到只需要处理两个链表的合并，之后两两合并，得到b1, b2 , … , 新链表，

接着继续两两合并得到c1, c2,… 直到最后只剩下一条有序链表。

这个方法在每一层的合并中，都需要处理所有的n个节点，而每次处理后链表的数量变为原来的一半。

总共要处理logk次，总的时间复杂度是O(n*logk).

使用的辅助空间大小为递归调用的深度，总共有logk层，所以空间复杂度为O(logk)
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

// 辅助函数，在还没有达到最基本情况前，不断递归调用自己
// 输入是链表数组，和当前要处理的在链表中开始和结束的下标，输出是合并后的链表
func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end { // 当开始下标等于结束下标时
		return lists[start] // 说明当前要处理的只有一个链表，直接返回它即可
	}
	if start > end { // 否则如果开始下标大于结束下标
		return nil // 无效的，直接返回空指针nil
	}
	// 如果不是上面两种情况，就分而治之
	// 先找到当前子数组的中间下标
	mid := start + (end-start)/2
	// 然后分别递归处理前一半和后一半链表
	left := merge(lists, start, mid)
	right := merge(lists, mid+1, end)
	// 得到的结果是两条合并后的有序链表
	// 最后再把这两条链表也合并即可
	return mergeTwoSortedLists(left, right)
}

func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1 // 以Next开始连接
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	// 连上l1剩余的链,连上l2剩余的链
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}
