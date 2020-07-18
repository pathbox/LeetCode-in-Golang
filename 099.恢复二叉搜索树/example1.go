package LeetCode099

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法一：最简单的方法是中序遍历一次然后排序再按照原来结构中序把数填回去
func inorderAppend(root *TreeNode, array *[]int) {
	if root == nil {
		return
	}
	inorderAppend(root.Left, array)
	*array = append(*array, root.Val)
	inorderAppend(root.Right, array)
}

func inorderFill(root *TreeNode, array []int, current *int) {
	if root == nil {
		return
	}
	inorderFill(root.Left, array, current)
	root.Val = array[*current]
	*current++
	inorderFill(root.Right, array, current)
}

func recoverTree(root *TreeNode) {
	array := []int{}
	inorderAppend(root, &array)
	sort.Ints(array)
	i := 0
	inorderFill(root, array, &i) // 将排完序的数组，重新中序遍历构造二叉搜索树
}
