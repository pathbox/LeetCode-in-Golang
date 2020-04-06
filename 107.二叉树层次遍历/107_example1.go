package LeetCode107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	level := 0
	if root == nil {
		return result
	}

	list := []*TreeNode{root}
	length := 1 //队列长度(即当前层节点数)
	for length > 0 {
		// 从队列中取出当前层
		for i := 0; i < length; i++ {
			//出队
			node := list[0]
			list = list[1:]

			//值放入result
			if len(result) > level {
				result[level] = append(result[level], node.Val)
			} else {
				result = append(result, []int{node.Val})
			}

			//下一层入队
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}
		length = len(list)
		level++
	}

	//数组翻转
	resultLength := len(result)
	left := 0
	right := resultLength - 1
	for left < right {
		temp := result[left]
		result[left] = result[right]
		result[right] = temp

		left++
		right--
	}
	return result
}
