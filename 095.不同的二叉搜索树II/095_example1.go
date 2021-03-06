package LeetCode095

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return dfs(1, n)
}

func dfs(start, end int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if start > end {
		res = append(res, nil)
		return res
	}

	for i := start; i <= end; i++ { // 循环start-end的整数
		leftTree := dfs(start, i-1)
		rightTree := dfs(i+1, end)

		for _, l := range leftTree {
			for _, r := range rightTree {
				cur := &TreeNode{
					Val: i,
				}
				cur.Left = l
				cur.Right = r
				res = append(res, cur)
			}
		}
	}
	return res
}
