package LeetCode102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := [][]int{}
	num := []*TreeNode{root} //树的节点队列
	for len(num) != 0 {
		length := len(num) // 每一层的节点个数
		result := []int{}  //用来存储每一层的值，到新的一层时需要进行清空
		for i := 0; i < length; i++ {
			result = append(result, num[i].Val)
			if num[i].Left != nil {
				num = append(num, num[i].Left)
			}
			if num[i].Right != nil {
				num = append(num, num[i].Right)
			}
		}
		ret = append(ret, result)
		num = num[length:]
	}
	return ret
}
