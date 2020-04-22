package LeetCode199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs右节点先遍历 遍历记录下每层的最后一个元素，如果同一层有多个元素，取最右边的元素

var res []int

func rightSideView(root *TreeNode) []int {
	res = make([]int, 0)
	dfs(root, 1) // 第一层开始
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if level > len(res) { // 如果 level < len(res)说明res中已经获取到了第一次(很有可能是右侧节点)遍历到的节点，则后面再到level这一层的时候，是不需要处理，进入res
		res = append(res, root.Val)
	}
	dfs(root.Right, level+1) // 右节点先遍历
	dfs(root.Left, level+1)
}
