package LeetCode098

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack = []*TreeNode{root} // 初始化时候带上根节点
	var minS = []int{-1 << 63}
	var maxS = []int{1<<63 - 1}
	for len(stack) > 0 {
		// pop
		pre := len(stack) - 1
		root, min, max := stack[pre], minS[pre], maxS[pre]
		stack, minS, maxS := stack[:pre], minS[:pre], maxS[:pre] // 出栈

		//push
		for root != nil {
			if root.Val <= min || max <= root.Val {
				return false // 不在min 和max 界限内
			}
			minS = append(minS, root.Val)
			maxS = append(maxS, max)
			stack = append(stack, root.Right) // 右节点入栈，先遍历的左子树
			max = root.Val
			root = root.Left
		}
	}
	return true
}
