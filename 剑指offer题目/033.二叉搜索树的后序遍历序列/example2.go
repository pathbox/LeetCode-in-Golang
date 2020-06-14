package Offer033

func verifyPostorder(postorder []int) bool {
	// 递归分治
	// 后序遍历性质： [ 左子树 | 右子树 | 根节点 ] ，即遍历顺序为 “左、右、根” 。
	// 二叉搜索树性质： 左子树任意节点的值 < 根节点的值；右子树任意节点的值 > 根节点的值。
	return recur33(postorder, 0, len(postorder)-1)
}

func recur33(postorder []int, left, right int) bool {
	// 终止条件： 当 left≥right ，说明子树节点少于或等于 1 个，无需判别正确性，因此直接返回 true ；
	if left >= right {
		return true
	}
	l := left
	for postorder[l] < postorder[right] {
		l++
	}
	middle := l // 记录中间位置

	for postorder[l] > postorder[right] {
		l++
	}
	var ret bool
	if l == right { // 满足二叉搜索树的性质 左子树|右子树|根节点 && 左 < 根 < 右
		ret = true
	} else {
		ret = false
	}

	return ret && recur33(postorder, left, middle-1) && recur33(postorder, middle, right-1) // 递归判断
}
