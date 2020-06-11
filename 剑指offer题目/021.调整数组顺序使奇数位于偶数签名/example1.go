package Offer021

// 首尾双指针

func exchange(nums []int) []int {
	head := 0
	tail := len(nums) - 1

	for head < tail {
		if nums[head]%2 == 0 && nums[tail]%2 != 0 {
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
			tail--
		} else {
			if nums[tail]%2 == 0 { // 如果尾部是偶数不断往前，直到是奇数就先停在那
				tail--
			}
			if nums[head]%2 != 0 { // 如果头部是奇数不断往后，直到是奇数就先停在那
				head++
			}
		}
	}
	return nums
}
