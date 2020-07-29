package LeetCode987

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 方法一：记录坐标
明显的，该解决方案有两个步骤：首先，找出每个节点所在的坐标，然后报告他们的坐标。

算法：

我们可以使用深度优先搜索找到每个节点的坐标。保持当前节点 (x, y)，移动的过程中，坐标变化为 (x-1, y+1) 或 (x+1, y+1) 取决于是左孩子还是右孩子。

我们通过 x 坐标排序，再根据 y 坐标排序，这样确保以正确的顺序添加到答案中
*/

type (
	pos struct {
		x, y, val int
	}
	result []pos
)

func (r result) Len() int {
	return len(r)
}

func (r result) Less(i, j int) bool {
	if r[i].x == r[j].x { // 同一列
		if r[i].y == r[j].y { // 同一列且同一行,数值小的放前面
			return r[i].val < r[j].val
		}
		return r[i].y < r[j].y // 同一列不同行,行数小的放前面
	}
	return r[i].x < r[j].x // 不同列则列数小的放前面
}

func (r result) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var loc result
	tree(root, 0, 0, &loc)
	sort.Sort(loc)
	var (
		index = 0
		x     = loc[index].x
		res   = [][]int{{loc[index].val}}
	)
	for i := 1; i < len(loc); i++ {
		if x != loc[i].x { // 不同列则赋值x,且结果新开一个数组
			x = loc[i].x
			index++
			res = append(res, []int{})
		}
		res[index] = append(res[index], loc[i].val)
	}
	return res
}

// 遍历数，标记每个节点的坐标值
func tree(root *TreeNode, x, y int, loc *result) {
	if root != nil { // 将坐标和值存入数组
		*loc = append(*loc, pos{x, y, root.Val})
		tree(root.Left, x-1, y+1, loc)
		tree(root.Right, x+1, y+1, loc)
	}
}
