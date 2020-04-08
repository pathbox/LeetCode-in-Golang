package LeetCode112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*时间复杂度：我们访问每个节点一次，时间复杂度为 O(N)O(N) ，其中 NN 是节点个数。
空间复杂度：最坏情况下，整棵树是非平衡的，例如每个节点都只有一个孩子，递归会调用 NN 次（树的高度），因此栈的空间开销是 O(N)O(N) 。但在最好情况下，树是完全平衡的，高度只有 \log(N)log(N)，因此在这种情况下空间复杂度只有 O(\log(N))O(log(N))
*/

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	if root.Left == nil && root.Right == nil { // root为叶子节点
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum) // 从左子树往下或右子树往下
}
