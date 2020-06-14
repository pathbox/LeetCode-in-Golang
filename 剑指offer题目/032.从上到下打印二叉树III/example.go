package Offer032

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1设置一个flag 标签判断是从左往右 还是从右往左
2 广度搜索采用队列的方式，首先是根节点作为首层入队 ，然后将根节点的左右节点入队，
3 当当请层遍历结束时 将当前层出队queue = queue[size:]
4 同时需要修改一下flag标签值isReverse = !isReverse
*/

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	var isReverse bool
	for len(queue) > 0 {
		length := len(queue)
		list := make([]int, length)
		for i := 0; i < length; i++ {
			node := queue[i]
			if !isReverse {
				list[i] = node.Val
			} else {
				list[length-i-1] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
		isReverse = !isReverse
		res = append(res, list)
	}
	return res
}
