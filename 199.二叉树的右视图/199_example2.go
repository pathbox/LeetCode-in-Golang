package LeetCode199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs 我们对树进行深度优先搜索，在搜索过程中，我们总是先访问右子树。那么对于每一层来说，我们在这层见到的第一个结点一定是最右边的结点
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			length--
			if len(res) == level { // 当数组长度等于当前 深度 时, 把当前的值加入数组，先入的右节点，每一层的第一个元素
				res = append(res, queue[0].Val)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}

			queue = queue[1:]
		}
		level++
	}
	return res
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		length := len(queue)
		size := length
		// 当前层的最后一个元素
		for i := 0; i < size; i++ {
			if i == size-1 {
				res = append(res, queue[0].Val)
			}
			// 将当前节点的非空左右子树入队列
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
