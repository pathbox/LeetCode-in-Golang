package LeetCode099

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
BST中序遍历 时间O(N) 空间O(N)
按中序遍历树。遍历后的数组应该是几乎排序的列表，其中只有两个元素被交换。
在线性时间内确定几乎排序数组中的两个交换元素 x 和 y。
再次遍历树，将值 x 的节点改为 y，将值 y 的节点改为 x
升序数组 ，只要我们找出不满足升序条件的数即可
*/

var last, first, second *TreeNode

func recoverTree(root *TreeNode) {
	last, first, second = nil, nil, nil
	dfs(root)
	first.Val, second.Val = second.Val, first.Val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if last != nil && root.Val <= last.Val {
		if first != nil {
			second = root
			return //剪枝
		}
		first, second = last, root
	}
	last = root
	dfs(root.Right)
}
