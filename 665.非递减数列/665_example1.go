package LeetCode665

func checkPossibility(nums []int) bool {
	count := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			tmp := nums[i]
			//先把nums[i]的值换成 前序数列的 最大值
			if i >= 1 {
				nums[i] = nums[i-1]
			} else {
				nums[i] = nums[i+1]
			}
			//这时说明前序数列的nums[i-1]也要大于 nums[i+1],所以说明改变 nums[i]没用，要改变nums[i+1]的值
			if nums[i] > nums[i+1] {
				nums[i] = tmp       //把nums[i]还原，也就值不改变nums[i]的值
				nums[i+1] = nums[i] //改变 nums[i+1]的值
			}
			count++
			//这样就确保 从0到 nums[i+1]是非递减数列 但是 这时候count加了两次 标示需要至少2次
		}
		if count == 2 {
			return false
		}
	}
	return true
}

// 借鉴大佬的解法，遍历数组，当找到nums[i]>nums[i+1]时flag记录一次，同时调整一下nums[i]，使当前下标为0到i的数组非递减。若第二次再找到nums[i]>nums，说明一次变换无法使真个数组非递减，返回false。
