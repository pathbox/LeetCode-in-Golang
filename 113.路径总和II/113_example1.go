package LeetCode113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	res = [][]int{}
	dfs(root, sum, []int{})
	return res
}

func dfs(root *TreeNode, sum int, stack []int) {
	if root == nil {
		return
	}
	stack = append(stack, root.Val)            // 遍历到某节点，Val入stack
	if root.Left == nil && root.Right == nil { // 判断是不是遍历到了叶子节点，是叶子节点，则判断sum值和Val值是否相等，sum值其实是每次遍历被减去后的值，相等说明找到了一条path
		if sum == root.Val {
			r := make([]int, len(stack))
			copy(r, stack)
			res = append(res, r)
		}
	}
	dfs(root.Left, sum-root.Val, stack)
	dfs(root.Right, sum-root.Val, stack)
}
