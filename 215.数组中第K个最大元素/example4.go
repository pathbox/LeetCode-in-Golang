package LeetCode215

func findKthLargest(nums []int, k int) int {
	nums = heapSort(nums)
	return nums[k-1]
}

func heapSort(nums []int) []int {
	lens := len(nums)
	// 建堆	lens/2后面都是叶子节点，不需要调整down()
	for i := lens / 2; i >= 0; i-- {
		down(nums, i, lens)
	}
	//将小根堆堆顶排到切片末尾(降序)
	for i := lens - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		lens--
		downs(nums, 0, lens)
	}
	return nums
}

func down(nums []int, i, lens int) { //小根堆
	min := i                    // i 父节点
	left, right := 2*i+1, 2*i+2 // 左右孩子
	if left < lens && nums[left] < nums[min] {
		min = left
	}
	if right < lens && nums[right] < nums[min] {
		min = right
	}
	if min != i {
		nums[min], nums[i] = nums[i], nums[min]
		down(nums, min, lens)
	}
}
