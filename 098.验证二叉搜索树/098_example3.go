package LeetCode098

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var last *TreeNode

/*
中序递归
递归方向 左->根->右
上一个节点 last 赋值时机
last.Val 必须小于 root.Val （因为中序遍历=升序遍历）
*/
func isValidBST(root *TreeNode) bool {
	last = &TreeNode{Val: -1 << 63}
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !dfs(root.Left) || root.Val <= last.Val { // 上一个节点的值应该要是更小的
		return false
	}
	last = root
	return dfs(root.Right)
}
