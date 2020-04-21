package LeetCode106

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	dfs(root)
	return root
}

// 左节点不断往右，右节点不断往左，像拉链一样拉紧
func dfs(root *Node) {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right
	for left != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}
	//递归的调用左右节点，完成同样的纵深串联, 树的纵深遍历
	dfs(root.Left)
	dfs(root.Right)
}

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/solution/dong-hua-yan-shi-san-chong-shi-xian-116-tian-chong/
