package Offer033

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}

	// 在后序遍历中，最后一个元素是根元素
	return juage(postorder, 0, len(postorder)-1)
}

func juage(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}
	var i int
	//从前面开始遍历，小于的当前根元素的值是左子树的，当找到第一个大于当前根元素的值，可以确定后半段的元素都应是在当前节点的右子树
	// end 是当前根节点
	for i := start; i < end; i++ {
		if postorder[i] > postorder[end] {
			break
		}
	}
	//如果后半段（右子树）里面有小于根元素的值的元素，就说明这个不是二叉搜索树的后序遍历，return false
	for j := i; j < end; j++ {
		if postorder[j] < postorder[end] {
			return false
		}
	}
	//递归检查左子树和右子树部分
	return juage(postorder, start, i-1) && juage(postorder, i, end-1)
}
