package LeetCode213

/*
环状排列意味着第一个房子和最后一个房子中只能选择一个偷窃
打劫第一家的话，最后一家就不能选，就是最后结果取dp1[n-2] （从0开始）
不打劫第一家的话，最后一家就可以选了，只需要将dp2[0]赋值为0，之后按照正常操作，最后取dp2[n-1]
最后把这两个值取最大值就是最终结果
*/
func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// dp[n] = max(dp[n-1], dp[n−2]+num)
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	dp1 := make([]int, n)
	dp2 := make([]int, n)
	dp1[0] = nums[0]               // 偷第一个房子
	dp2[0] = 0                     // 不偷第一个房子
	dp1[1] = Max(nums[0], nums[1]) // 偷第一个还是第二个
	dp2[1] = nums[1]               // 偷第二个
	for i := 2; i < n; i++ {
		dp1[i] = Max(dp1[i-1], dp1[i-2]+nums[i])
		dp2[i] = Max(dp2[i-1], dp2[i-2]+nums[i])
	}
	return Max(dp1[n-2], dp2[n-1]) // dp1[n-2]：打劫第一家不能打劫最后一家，dp2[n-1]：不打劫第一家能打劫最后一家
}

//------------------------------
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	return max(djjs(nums[1:]), djjs(nums[:len(nums)-1]))
}

func djjs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		tmp := second
		second = max(nums[i]+first, second)
		first = tmp
	}

	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
class Solution:
	def rob(self, nums: [int]) -> int:
		cur, pre = 0, 0
		for num in nums:
			cur, pre = max(pre + num, cur), cur
		return cur
	return max(my_rob(nums[:-1]),my_rob(nums[1:])) if len(nums) != 1 else nums[0]
*/
