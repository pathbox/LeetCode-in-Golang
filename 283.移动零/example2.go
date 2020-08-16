package LeetCode283

func moveZeroes(nums []int) {
	if nums == nil {
		return
	}

	// j 为慢指针，永远指向第一个0，用来和i指向的非0数字交换
	var j int
	for i := 0; i < len(nums); i++ {
		// 如果i也遇到了0，那么直接跳过，直到遇见第一个非0的数字
		if nums[i] == 0 {
			continue
		}
		// 然后和j指向的0互换
		nums[i], nums[j] = nums[j], nums[i]
		// 接着j指向的就不是非0的数字了，所以往后移动一位
		j++
	}
}
