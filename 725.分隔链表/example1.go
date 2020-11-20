package LeetCode725

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
巧妙类比：
（1）先数出有多少块小糖，得n块；
（2）将n块小糖分给k个小朋友，从左往右均分(n/k)块；
（3）将剩下的(n%k)块小糖再从左往右1人1块，直至分完
*/

// O(N+K) O(K)
func splitListToParts(root *ListNode, k int) []*ListNode {
	n := 0
	p := root
	for p != nil {
		n++
		p = p.Next
	}
	ans := make([]*ListNode, k) // 没有赋值到的节点会是nil
	p = root
	if n <= k { // 1. n 小于等于k的情况
		i := 0
		for p != nil {
			ans[i] = p
			p = p.Next
			ans[i].Next = nil
			i++
		}
		return ans
	}

	width := n / k      // 平均分
	left := n - width*k // 平均分完后还剩几个
	for i := 0; i < k; i++ {
		ans[i] = p // 加入本部分的头结点
		for j := 1; j < width; j++ {
			p = p.Next
		}
		if left > 0 { // 剩下的分配可以在一个新的循环中，也可以在同一个循环中
			left--
			p = p.Next
		}
		next := p.Next
		p.Next = nil // 这一部分结束
		p = next     // 下一部分
	}
	return ans
}
