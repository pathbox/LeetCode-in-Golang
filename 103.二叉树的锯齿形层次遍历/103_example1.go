package LeetCode102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res = [][]int{}
	var queue = []*TreeNode{root}
	var level, counter, knife, cur int
	for len(queue) > 0 {
		counter = len(queue)
		knife = counter
		res = append(res, []int{})
		for counter > 0 {
			// BFS入队列 每层的队列都是先左后右
			cur = knife - counter
			if queue[cur].Left != nil {
				queue = append(queue, queue[cur].Left)
			}
			if queue[cur].Right != nil {
				queue = append(queue, queue[cur].Right)
			}

			counter--

			if level%2 != 0 { // level 是偶数 cur，从队列尾部取，表示从右边Z字形遍历，level是奇数，cur是从头部小数索引位置取
				cur = counter
			}
			res[level] = append(res[level], queue[cur].Val)
		}
		queue = queue[knife:] // 清空这一层队列数据，为下一层准备
		level++
	}
	return res
}
