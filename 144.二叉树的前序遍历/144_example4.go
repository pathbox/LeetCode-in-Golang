package LeetCode144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var max *TreeNode
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil && max.Right != root {
				max = max.Right
			}

			if max.Right == nil {
				res = append(res, root.Val)
				max.Right = root.Right
				root = root.Left
			} else {
				root = root.Right
				max.Right = nil
			}
		}
	}
	return res
}

/*
https://leetcode-cn.com/problems/binary-tree-preorder-traversal/solution/leetcodesuan-fa-xiu-lian-dong-hua-yan-shi-xbian-2/

Morris的整体思路就是将 <u>以某个根结点开始，找到它左子树的最右侧节点之后与这个根结点进行连接</u>
我们可以从 图2 看到，如果这么连接之后，cur 这个指针是可以完整的从一个节点顺着下一个节点遍历，将整棵树遍历完毕，直到 7 这个节点右侧没有指向
*/
