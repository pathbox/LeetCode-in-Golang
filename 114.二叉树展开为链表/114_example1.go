package LeetCode114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	for root != nil {
		//左子树为 null，直接考虑下一个节点
		if root.Left == nil {
			root = root.Right
		} else {
			// 找左子树最右边的节点
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			//将原来的右子树接到左子树的最右边节点
			pre.Right = root.Right
			// 将左子树插入到右子树的地方
			root.Right = root.Left
			root.Left = nil
			// 考虑下一个节点
			root = root.Right
		}
	}
}

/*
空间复杂度是O(1),在原来树的结构上操作，而不是借用额外的存储空间

1. 将左子树插入到右子树的地方
2. 将原来的右子树接到左子树的最右边节点
3. 考虑新的右子树的根节点，一直重复上边的过程，直到新的右子树为 null
*/
