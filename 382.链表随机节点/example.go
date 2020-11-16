package LeetCode382

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	r    *rand.Rand
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	solution := Solution{
		head: head,
		r:    r,
	}
	return solution
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	i := 2
	cur := this.head.Next
	val := this.head.Val
	for cur != nil {
		if this.r.Intn(i)+1 == i {
			val = cur.Val
		}
		i++
		cur = cur.Next
	}
	return val
}

// ### 复杂度分析
// 时间复杂度：O（n）至少要遍历到n个元素
// 空间复杂度：O（1）常数空间复杂度

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
