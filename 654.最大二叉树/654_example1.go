package LeetCode654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 以数组最大值左右两边分别为左右子树继续递归
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	idx := maxIndex(nums)
	root := &TreeNode{
		Val:   nums[idx],
		Left:  constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
	}
	return root
}

func maxIndex(nums []int) int {
	idx := 0
	for i := range nums {
		if nums[idx] < nums[i] {
			idx = i
		}
	}
	return idx
}
