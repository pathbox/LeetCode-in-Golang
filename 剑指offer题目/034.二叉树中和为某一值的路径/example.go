package Offer034

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	dfs(root, sum, []int{}, &ret)
	return ret
}

func dfs(root *TreeNode, sum int, arr []int, ret *[][]int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)

	if root.Val == sum && root.Left == nil && root.Right == nil {
		//slice是一个指向底层的数组的指针结构体
		//因为是先序遍历，如果 root.Right != nil ,arr 切片底层的数组会被修改
		//所以这里需要 copy arr 到 tmp，再添加进 ret，防止 arr 底层数据修改带来的错误
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*ret = append(*ret, tmp)
	}

	dfs(root.Left, sum-root.Val, arr, ret)
	dfs(root.Right, sum-root.Val, arr, ret)

	arr = arr[:len(arr)-1]
}
