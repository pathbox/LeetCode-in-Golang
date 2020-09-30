package LeetCode199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leftSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			length--               // 每一层的第一个值
			if len(res) == level { // 当数组长度等于当前 深度 时, 把当前的值加入数组, 先入左节点，每一层的第一个元素判断
				res = append(res, queue[0].Val)
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		level++
	}
	return res
}
