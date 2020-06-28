package Offer061

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	sub := 0
	for i := 0; i < 4; i++ {
		// 大小王
		if nums[i] == 0 {
			continue
		}
		//如果扑克牌（非0数字）重复，不是顺子
		if nums[i] == nums[i+1] {
			return false
		}
		//差值记录
		sub += nums[i+1] - nums[i]
	}
	//总的差值小于5（或sub<=4）则为顺
	return sub < 5
}
