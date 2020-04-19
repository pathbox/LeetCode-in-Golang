package LeetCode106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	l := len(postorder) - 1
	for k := range inorder {
		if inorder[k] == postorder[l] { // postorder的最后一个值为每个子树的root
			return &TreeNode{
				Val:   inorder[k],
				Left:  buildTree(inorder[:k], postorder[:k]),
				Right: buildTree(inorder[k+1:], postorder[k:l]),
			}
		}
	}
	return nil

}

/*
通常从先序序列或者后序序列开始，根据不同遍历方法的规律，选择合适的节点构造树。
例如：先序序列的 第一个 节点是根节点，然后是它的左孩子，右孩子等等。
后序序列的 最后一个 节点是根节点，然后是它的右孩子，左孩子等等。

从先序/后序序列中找到根节点，根据根节点将中序序列分为左子树和右子树。从中序序列中获得的信息是：如果当前子树为空（返回 None），否则继续构造子树。

https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/cong-zhong-xu-yu-hou-xu-bian-li-xu-lie-gou-zao-e-5/

nice: https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/tu-jie-gou-zao-er-cha-shu-wei-wan-dai-xu-by-user72/
*/
