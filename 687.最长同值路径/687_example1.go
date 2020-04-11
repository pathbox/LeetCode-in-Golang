package LeetCode687

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	arrowLength(root, &ans)
	return ans
}

func arrowLength(root *TreeNode, num *int) int {
	if root == nil {
		return 0
	}
	l := arrowLength(root.Left, num)
	r := arrowLength(root.Right, num)
	var al, ar int
	if root.Left != nil && root.Left.Val == root.Val {
		al = l + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		ar += r + 1
	}
	*num = int(math.Max(float64(*num), float64(ar+al)))
	return int(math.Max(float64(al), float64(ar)))
}

/*
递归函数的逻辑
41. 当我们递归调用 node->left 和 node->right 时，就可以得到左/右子树与其数值相同节点的长度
42. 只需要判断左/右子树的数值与根节点的数值是否相同，不相同的话就舍弃，相同的话长度就加一
43. 那么与 node 值相同的路径长度就是左/右两条路中较长的

递归函数只能取得 node 的同值路径长度，这里还有一些其他情况需要考虑
51. 如果最长的路径是由一个其他数值的子节点开始向下出发的情况，在递归中间因为数值不同被舍弃了
52. 是由一个子节点，向左右两个方向同时向下出发的情况（情况 1）

在递归函数的参数中使用一个引用变量，记录全局最大值
61. 当计算出一个节点的返回值时，将以这个节点，连接左右两边的路径长度，更新到全局最大值
62. 然后再返回

这个全局最大值包含了以上所有情况，就是题目的答案

*/
