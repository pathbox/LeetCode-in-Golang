package LeetCode222

* type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
利用题目中 完全二叉树的特性
1.先求出左右子树的高度left, right
2.若左右子树高度相等，则左子树一定是满二叉树，其节点数 = 2^left - 1，加上root节点 = 2^left
3.若左右子树高度不等，则右子树一定是满二叉树，其节点数 = 2^right - 1，加上root节点 = 2^right
4.计算满二叉树的节点数，可使用位运算，算术左移
时间复杂度 O(logn * logn)  O(1)
*/

func countNodes(root *TreeNode) int { // countNodes方法本身就是计算树的节点数的方法,所以可以用在26 28行
	if root == nil {
		return 0
	}
	// 先求出左右子树的高度left, right
	left := countLevel(root.Left)
	right := countLevel(root.Right)
	if left == right {
		return countNodes(root.Right) + (1<<left) // 递归计算右子树节点数+左子树节点数
	} else {
		return countNodes(root.Left) + (1<<right)
	}
}
// 计算节点数时也可以用这个方法代替
// func countMyNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return 1 + countMyNodes(root.Left) + countMyNodes(root.Right)
// }

// 求树的高度
func countLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(countLevel(root.Left), countLevel(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}