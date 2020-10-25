package LeetCode651

func maxA(n int) int {
	return dp(n, 0, 0) // 初始状态
}

/*
在屏幕上的数量、在剪切板的数量是状态
A 、 C-A 、 C-C 、 C-V （ Ctrl 简写为C）是选择
*/

// n表示还剩的按键次数  num表示在屏幕上的字母数量，copy表示在剪切板上的字母数量
func dp(n, num, copy int) int {
	if n <= 0 { // 次数用完了， 返回此时屏幕上的字母数量
		return num
	}

	return max([]int{
		dp(n-1, num+1, copy),    // 按A键
		dp(n-1, num+copy, copy), // 粘贴键
		dp(n-2, num, num),       // 全选+复制
	})

}

func max(nums []int) int {
	m := nums[0]
	for v := range nums {
		if m < v {
			m = v
		}
	}
	return m
}
