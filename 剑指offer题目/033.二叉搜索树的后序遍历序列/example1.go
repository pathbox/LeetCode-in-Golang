package Offer033

/*
划分左子树,我们能知道现在这个树的根节点为[5],从头开始遍历节点,直到我们遇到比根节点大的或者根节点,之前那些比根节点小的值是左子树的节点,此树的左子树序列为[1]
划分右子树,可以看到当遍历到6的时候,6>5,按照上面总结的规律,说明6之后到根节点的序列都应该是右子树,即[6,3,2],这是就需要判断下是否在右子树的序列中存在比根节点小的元素,如果有的话就说明是不符合条件的,显然[3,2]都是比[5]要小的,不符合条件
如果不存在在右子树中比根节点小的元素,递归的遍历上述左右子树

*/

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}
	//在后序遍历中，最后一个元素是根元素
	return juage(postorder, 0, len(postorder)-1)
}

func juage(postorder []int, start, end int) bool {
	// 返回条件
	if start >= end {
		return true
	}
	var i int
	//从前面开始遍历，小于的当前根元素的值是左子树的，当找到第一个大于当前根元素的值，可以确定后半段的元素都应是在当前节点的右子树
	for i = start; i < end; i++ {
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
	//递归检查左子树和右子树部分 postorder[end]已经不需要继续检查
	return juage(postorder, start, i-1) && juage(postorder, i, end-1)
}
