package LeetCode437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var s int

func pathSum(root *TreeNode, sum int) int {
	res = 0
	s = sum
	run(root)
	return res
}

/*
	两层递归，不断的递归遍历树，遍历某个节点的时候，再以这个节点为root节点，遍历其左右子树，求和寻找满足的情况
*/
func run(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root, 0) // 前序遍历
	run(root.Left)
	run(root.Right)

}

// 继续的dfs 前序遍历
func dfs(root *TreeNode, cur int) {
	if root != nil {
		cur += root.Val // 处理root，然后继续遍历左右子树
		if s == cur {
			res++
		}
		dfs(root.Left, cur)
		dfs(root.Right, cur)
	}
}
