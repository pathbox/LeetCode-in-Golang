package LeetCode1024

func videoStitching(clips [][]int, T int) int {
	quickSort2D(clips)
	if len(clips) == 0 {
		return -1
	}
	result := genJumpRecord(clips, T)

	return result
}

func genJumpRecord(diffJumps [][]int, T int) int {
	result := 0
	preEnd := 0
	curIndex := 0
	curStart := diffJumps[0][0]
	maxEnd := preEnd
	for curIndex < len(diffJumps) && curStart <= preEnd {
		for curStart <= preEnd && curIndex < len(diffJumps) { // 在下次起始点不超过本次范围的情况下，寻找一个覆盖范围最广的片段.下一个片段的start要在上一个片段的区间内，才有交集，才能连接上
			if diffJumps[curIndex][1] > maxEnd {
				maxEnd = diffJumps[curIndex][1]
			}
			curIndex++
			if curIndex < len(diffJumps) {
				curStart = diffJumps[curIndex][0]
			}
		}
		result++
		preEnd = maxEnd
		if maxEnd >= T {
			return result
		}
	}
	return -1
}

func quickSort2D(nums [][]int) {
	if len(nums) == 0 {
		return
	}
	left := 0
	right := len(nums) - 1
	pivot := nums[left]
	for left < right {
		for nums[right][0] >= pivot[0] && left < right {
			right--
		}
		nums[left] = nums[right]
		for nums[left][0] < pivot[0] && left < right {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	if left > 0 {
		quickSort2D(nums[:left])
	}
	if len(nums) > left+1 {
		quickSort2D(nums[left+1:])
	}
}
