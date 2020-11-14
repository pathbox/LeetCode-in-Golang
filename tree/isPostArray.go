
package tree

// arr 是后续遍历的二叉搜索树，判断是否为二叉搜索树 end位置是头结点
func isPostArray([]int arr) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	return isPost(arr, 0, len(arr)-1)
}

func isPost([]int arr, start, end int) bool {
	if start == end {
		return true
	}

	less := -1
	more := end
	for i := start; i < end; i++{
		if arr[i] < arr[end] {
			less = i
		} else {
			if more == end {
				more = i
			}
		}
	}
	if less == -1 || more == end { // 表示start->end-1 整个是右子树或左子树,所以less 或 more的其中一个值没改变
		return isPost(arr, start, end -1)
	}
	if less != more-1{
		return false
	}
	return isPost(arr, start, less) && isPost(arr, more, end-1)
}

func posToBST([]int posArr, start, end int) *Node {
	if start > end {
		return nil
	}
	head := &Node{Val: posArr[end]}
	less := -1
	more := end
	for i := start; i < end; i++ {
		if posArr[end] > posArr[i] {
			less = i
		} else if more == end {
			more = i
		}
	}
	head.Left = posToBST(posArr, start, less)
	head.Right = posToBST(posArr, more, end-1)
}
