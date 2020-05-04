package LeetCode1104

func pathInZigZagTree(label int) []int {
	// 构造树的深度
	curMax, tree := 0, []int{}
	for label > curMax {
		curMax = bpow(len(tree)+1) - 1
		tree = append(tree, curMax)
	}
	// 从底层往上遍历
	curLevel := len(tree) // 树的深度
	curPos := curLevel    // 当前遍历的层数
	curLable := label     // 当前元素

	curOffset := 0 // 当前层相对首个元素的偏移量
	if curLevel%2 == 0 {
		curOffset = bpow(curPos) - label // 偶数行从大到小
	} else {
		curOffset = label - bpow(curPos-1) + 1 // 奇数行从小到大
	}

	result := make([]int, curLevel)
	result[curLevel-1] = curLable
	for i := curLevel - 2; i >= 0; i-- {
		curPos--
		if curOffset%2 == 0 {
			curOffset = curOffset / 2
		} else {
			curOffset = curOffset/2 + 1
		}

		if curPos%2 == 0 {
			curLable = bpow(curPos) - curOffset
		} else {
			curLable = bpow(curPos-1) + curOffset - 1
		}
		result[i] = curLable
	}
	return result
}

func bpow(y int) int {
	if y <= 0 {
		return 1
	}
	if y == 1 {
		return 2
	}
	return 2 << uint(y-1)
}
