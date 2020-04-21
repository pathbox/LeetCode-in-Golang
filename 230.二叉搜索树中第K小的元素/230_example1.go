package LeetCode230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 可以使用 BST 的特性：BST 的中序遍历是升序序列
// 通过构造 BST 的中序遍历序列，则第 k-1 个元素就是第 k 小的元素
// 时间复杂度：O(N)O(N)，遍历了整个树。
// 空间复杂度：O(N)O(N)，用了一个数组存储中序序列
func kthSmallest(root *TreeNode, k int) int {
	var res []int
	inOrder(root, &res)
	return res[k-1]
}

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}
