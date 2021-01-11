package LeetCode872

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	result1 := []int{}
	result2 := []int{}
	dfs(root1, &result1)
	dfs(root2, &result2)

	return ifEqual(result1, result2)
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*result = append(*result, root.Val)
	}
	dfs(root.Left, result)
	dfs(root.Right, result)
}

func ifEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
