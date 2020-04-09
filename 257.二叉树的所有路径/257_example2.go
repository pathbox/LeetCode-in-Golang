package LeetCode257

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := TreePaths(root, "", []string{})
	return ans
}

func TreePaths(cur *TreeNode, s string, ans []string) []string {
	if cur == nil {
		return ans
	}
	s += fmt.Sprintf("%d->", cur.Val)        // 累加当前路径
	if cur.Left == nil && cur.Right == nil { // 如果是叶子节点，则这条路径到最后了，返回
		ans = append(ans, s[:len(s)-2])
		return ans
	}
	ans = TreePaths(cur.Left, s, ans)
	ans = TreePaths(cur.Right, s, ans)
	return ans
}
