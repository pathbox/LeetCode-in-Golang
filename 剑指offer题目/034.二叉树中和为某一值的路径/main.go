package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	dfs(root, sum, []int, &ret)
	return ret
}

func main() {

}
