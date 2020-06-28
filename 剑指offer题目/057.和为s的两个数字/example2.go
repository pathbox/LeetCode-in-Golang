package Offer057

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return nil
	}
	left := 0
	right := len(nums) - 1
	//大于目标值的数字全都排除
	for nums[right] > target {
		right--
	}
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		}
		if nums[left] > target-nums[right] {
			right--
		} else {
			left++
		}
		//同理有，这里选择一种情况写即可，否则会出现错误
		/*
		   if nums[right]>target-nums[left]{
		       right--
		   }else{
		       left++
		   }
		*/
	}
	return nil
}
