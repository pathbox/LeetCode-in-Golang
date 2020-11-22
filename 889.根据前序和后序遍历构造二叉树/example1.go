package LeetCode889

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/solution/kan-wo-jiu-gou-liao-san-chong-bian-li-fang-shi-gou/
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/solution/tu-jie-889-gen-ju-qian-xu-he-hou-xu-bian-li-gou-2/
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if pre == nil || len(pre) == 0 {
		return nil
	}
	if post == nil || len(post) == 0 {
		return nil
	}

	root := &TreeNode{Val: pre[0]}
	if len(pre) == 1 {
		return root
	}
	n := len(pre)
	for i := 0; i < len(post); i++ {
		if pre[1] == post[i] { // pre的下一位pre[1]去post定位左子树
			leftCount := i + 1 // 需要+1 因为是右开区间
			root.Left = constructFromPrePost(pre[1:leftCount+1], post[0:leftCount])
			root.Right = constructFromPrePost(pre[leftCount+1:n], post[leftCount:n-1])
			break
		}
	}

	return root
}
