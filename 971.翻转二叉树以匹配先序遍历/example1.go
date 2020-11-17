package LeetCode971

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度：O(N)，其中 NN 是给定树中节点的数量。

// 空间复杂度：O(N)
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var res []int
	dfs(root, &voyage, &res)
	return res
}

func dfs(root *TreeNode, voyage, res *[]int) bool {
	if root == nil {
		return true
	}
	// 根节点的值和voyage第0个元素不符，无法通过反转节点使使先序遍历符合voyage
	if root.Val != (*voyage)[0] {
		*res = []int{-1}
		return false
	}
	// 当且仅当根节点的左右节点都不为空，且右儿子的值等于voyage的第1个元素，需要反转根节点的左右儿子
	if root.Left != nil && root.Right != nil && root.Right.Val == (*voyage)[0] {
		*res = append(*res, root.Val)
		root.Left, root.Right = root.Right, root.Left
	}
	// 消耗掉voyage第0个元素
	*voyage = (*voyage)[1:]
	// 如果对左儿子递归的结果为false，就进行剪枝操作
	ok := dfs(root.Left, voyage, res)
	if !ok {
		return false
	}
	return dfs(root.Right, voyage, res)
}
