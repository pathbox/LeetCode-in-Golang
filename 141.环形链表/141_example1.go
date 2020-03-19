package LeetCode021

type ListNode struct {
	Val  int
	Next *ListNode
}

// 开一个哈希表记录该节点是否已经遍历过，值记录节点索引
// 该节点遍历过，形成了环
// 将遍历过的节点记录下来，如果又遍历到了，表示链表有环，时间复杂度O(n)，空间复杂度O(n)
func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = struct{}{}
		head = head.Next // 链表头指针不断移动
	}
	return false
}
