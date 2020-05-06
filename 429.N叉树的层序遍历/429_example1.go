package LeetCode429

type Node struct {
	Val      int
	Children []*Node
}

var res [][]int

func levelOrder(root *Node) [][]int {
	res = [][]int{}
	if root == nil {
		return res
	}
	queue := []*Node{root} // 第一层的队列就是root
	var level int          // 定义一个level
	for len(queue) > 0 {
		counter := len(queue) // 当前这层level队列中所有元素数量
		res = append(res, []int{})
		for i := 0; i < counter; i++ {
			if queue[i] != nil {
				res[level] = append(res[level], queue[i].Val)
				for _, n := range queue[i].Children {
					queue = append(queue, n)
				}
			}
		}
		queue = queue[counter:] // [0:counter]已经遍历过
		level++
	}
	return res
}

/*
我们可以使用递归来解决这个问题，通常我们不能使用递归进行广度优先搜索。这是因为广度优先搜索基于队列，而递归运行时使用堆栈，适合深度优先搜索

dfs： 栈
bfs： 队列
*/
