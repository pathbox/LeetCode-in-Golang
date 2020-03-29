package LeetCode617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代法
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	// 基于t1来构建返回结果列表
	stack := make([][]*TreeNode, 0)
	stack = append(stack, []*TreeNode{t1, t2})
	for len(stack) != 0 {
		// pop操作
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 有一个为nil，表示该分支已走完，可以结束处理
		if t[0] == nil || t[1] == nil {
			continue
		}

		// 处理根节点
		t[0].Val += t[1].Val

		// 处理左子树
		if t[0].Left == nil {
			t[0].Left = t[1].Left
		} else {
			stack = append(stack, []*TreeNode{t[0].Left, t[1].Left})
		}
		// 处理右子树
		if t[0].Right == nil {
			t[0].Right = t[1].Right
		} else {
			stack = append(stack, []*TreeNode{t[0].Right, t[1].Right})
		}
	}
	return t1
}
