package LeetCode189

func rotate(nums []int, k int) {
	k %= len(nums) // 是个环移动 所以用%取余
	ans := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, ans) // 赋值回给nums
}

/*
3次反转
原始数组                  : 1 2 3 4 5 6 7
反转所有数字后             : 7 6 5 4 3 2 1
反转前 k 个数字后          : 5 6 7 4 3 2 1
反转后 n-k 个数字后        : 5 6 7 1 2 3 4 --> 结果
*/
