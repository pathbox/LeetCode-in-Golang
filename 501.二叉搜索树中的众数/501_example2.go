package LeetCode501

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	nums := []int{}
	cur, max, pre := 0, 0, &TreeNode{}
	helper(root, &nums, &cur, &max, &pre)
	res := make([]int, len(nums))
	for i, v := range nums {
		res[i] = v
	}
	return res
}

// 中序遍历
func helper(root *TreeNode, nums *[]int, cur, max *int, pre **TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left, nums, cur, max, pre) // 先遍历左子树
	if (*pre) != nil {
		if root.Val == (*pre).Val {
			*cur++
		} else {
			*cur = 1
		}
	}
	if *cur > *max {
		*max = *cur
		*nums = []int{} // 重新生成nums数组
		*nums = append(*nums, root.Val)
	} else if *cur == *max { // 把root.Val值接入到当前数组中，这样 可能会有多个众数，而不是唯一众数
		*nums = append(*nums, root.Val)
	}
	(*pre) = root // 当前节点赋值给pre前一个指针
	helper(root.Right, nums, cur, max, pre)
}
