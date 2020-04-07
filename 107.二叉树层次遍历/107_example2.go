package LeetCode107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int
	var queue []TreeNode

	if root == nil {
		return ret
	}

	queue = append(queue, *root)
	for len(queue) > 0 {
		var tmp []int
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			tmp = append(tmp, queue[i].Val)

			if queue[i].Left != nil {
				queue = append(queue, *queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, *queue[i].Right)
			}
		}
		ret = append(ret, tmp)
		queue = queue[queueLen:]
	}

	lp := 0
	rp := len(ret) - 1
	for lp < rp {
		ret[lp], ret[rp] = ret[rp], ret[lp]
		lp++
		rp--
	}
	return ret
}

/*
结果需要自底向上，直接自底向上显然不好实现，所以可以采用自顶向下层次遍历，然后再反转即可得到。
如果用一个队列进行层次遍历，会有一个麻烦，就是不知道每一层什么时候结束，就没办法将每一层的数据分别用不同集合来存储。
所以可以采用两个队列来实现，一个队列遍历当前层，一个队列存储下一层，交替使用，达到目的

方法二: 深搜实现层次遍历
使用深搜来实现层次遍历的核心就是必须记录住当前是第几层，只有知道了当前是第几层才能往对应层的集合添加数据
*/
