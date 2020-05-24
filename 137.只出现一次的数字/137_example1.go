package LeetCode137

// https://leetcode-cn.com/problems/single-number-ii/solution/zhi-chu-xian-yi-ci-de-shu-zi-ii-by-leetcode/
// https://leetcode-cn.com/problems/single-number-ii/solution/single-number-ii-mo-ni-san-jin-zhi-fa-by-jin407891/

// https://leetcode-cn.com/problems/single-number-ii/solution/man-hua-gan-jiu-dui-liao-ao-li-gei-by-ivan1-4/

// 异或操作可以让两个相同的数归0

// 逐位判断
func singleNumber(nums []int) int {
	var res int32
	var i uint
	for ; i < 32; i++ {
		// 逐位考虑
		var cnt uint
		for k := 0; k < len(nums); k++ {
			cnt += (uint(nums[k]) >> i) & 1 // 统计每个元素这一位出现1的次数
		}
		if cnt%3 != 0 { // 这一位的1出现非3次
			res = res | (1 << i) // 把这一位记到结果的相应位，位数只有0/1
		}
	}
	return int(res) // 返回32位完整的结果
}
/*
func singleNumber(nums []int) int {
	number, res := 0,0
	for i :=0; i < 64; i++{
		number = 0
		for _, k := range nums {
			number += (k >> i) & 1
		}
		res |= (number) % 3 << i
	}
	return res
}

*、