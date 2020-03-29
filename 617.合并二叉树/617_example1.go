package LeetCode617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// 谁为nil，则返回另一方子树
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val // 前序遍历，先操作根节点，再遍历左子树和右子树
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// 时间复杂度：O(N)O(N)，其中 NN 是两棵树中节点个数的较小值。

// 空间复杂度：O(N)O(N)，在最坏情况下，会递归 NN 层，需要 O(N)O(N) 的栈空间。

// 树的节点并非只是一个单个节点，而是表示一个子树的节点，其下有可能还连这一个子树
// 关于递归代码的编写，以树的遍历为例子，为什么我一定要纠结代码是如何写的呢？实际上递归就是不断的入栈，达到返回条件，则出栈，返回到上一层，继续递归或返回。如果能够明白树是如何递归遍历的，纠结于思考代码是怎么递归的也许是一个死胡同。而从实际中明白起递归遍历的方式，才是更好的方式
