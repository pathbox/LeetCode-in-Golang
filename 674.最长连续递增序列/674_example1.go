package LeetCode674

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	ans, count := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] { // 不断的进行递增比较
			count++
		} else {
			count = 1
		}
		if count > ans { // 将更大的值赋值给ans,ans最后返回
			ans = count
		}
	}
	return ans
}
