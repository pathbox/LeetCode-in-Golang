package LeetCode655

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归计算树的高度是典型的递归运算 计算树的高度 递归方法
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printTree(root *TreeNode) [][]string {
	// 计算输出数组的规模
	h := getHeight(root)
	// 宽度（最底层的宽度）， 2^n - 1
	w := (1 << h) - 1
	// 构造答案二维数组（slice）
	ans := make([][]string, h)
	for i := range ans {
		ans[i] = make([]string, w)
	}
	fill(root, &ans, 0, 0, w-1)
	return ans
}

// fill 递归的填充。
//
// 对当前传入的数，如果是空树，则返回。
// 我们获得当前可填充范围，计算出中位数mid，在二维数组的[h][mid]填充节点值，
// 然后递归对左节点，右节点调用fill方法。
func fill(root *TreeNode, ans *[][]string, h, l, r int) {
	if root == nil {
		return
	}
	mid := (l + r) / 2 // mid列的位置
	(*ans)[h][mid] = strconv.Itoa(root.Val)
	fill(root.Left, ans, h+1, l, mid-1)
	fill(root.Right, ans, h+1, mid+1, r)
}
