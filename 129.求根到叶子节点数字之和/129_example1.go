package LeetCode114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 前序遍历求和
在 叶节点 的时候 对 res 做加法，
在下降到 子节点 的时候做乘法,每下降一层，*10
*/
var res int

func sumNumbers(root *TreeNode) int {
	res = 0
	if root != nil {
		dfs(root, 0)
	}
	return res
}

// 没有用数组，而是用了sum，每个递归的层都带进去计算
func dfs(root *TreeNode, sum int) {
	if root.Left == nil && root.Right == nil { // 到叶子节点
		res += sum*10 + root.Val
		return
	}
	if root.Left != nil {
		dfs(root.Left, sum*10+root.Val) //子节点 乘以10
	}
	if root.Right != nil {
		dfs(root.Right, sum*10+root.Val) //子节点 乘以10
	}
}
