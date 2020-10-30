package LeetCode477

// 将所有数转换的成二进制，现在我们先看第0位，求任意两个数第0位的汉明距离的总和。我们从左往右遍历数组，当遍历到第i个数nums[i]时，nums[i]第0位能贡献的汉明距离等于i左边所有数的第0位与nums[i]的第0位不同的数个数，本来还需要加上右边的数，但为了避免重复计算，我们只看左边的数，所以我们可以一边遍历一边统计出现的0和1的个数。这样当我们遍历完数组，就可以得到所有数的第0位能贡献的汉明距离。而题目的答案就是求出所有位能贡献的汉明距离

func totalHammingDistance(nums []int) int {
	count := 0
	for i := uint(0); i < 32; i++ {
		c, b := 0, 1<<i // b是第i位是1的掩码 c是数子位置为1的总和 i要是uint
		for _, num := range nums {
			if num&b != 0 {
				c++
			}
		}
		count += c * (len(nums) - c) // len(nums) - c: 是数子位置为0总和
	}
	return count
}
