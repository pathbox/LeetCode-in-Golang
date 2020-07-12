package LeetCode060

// 1. 将数字变成字节数组，如 n = 4 得到 ['1', '2', '3', '4'], 并计算 n! 值
// 2. 考虑排列的第一个字符，以 n = 4 为例
// - 以 1, 2, 3, 4为首的排列各占 4!/4 = 6 个
// - 通过对比 k 值，以 k = 9 为例，9 > 6 && 9 < 6*2 则可以确认第9个排列的第一个字符应该为 2
// - 确认之后我们将 2 字符从字节数组中去掉，则只剩下三个 1, 3, 4
// - 此时我们需要从剩下三个字符的排列中寻找第 9-6 = 3 个，新一轮循环开始
func getPermutation(n int, k int) string {
	nums, total := make([]byte, 0, n), 1
	for i := 1; i <= n; i++ {
		nums = append(nums, byte(i)+'0')
		total *= i
	}
	res := make([]byte, 0, n)
	for i := n; i > 0; i-- {
		t := total / i
		idx := k/t - 1
		if k%t > 0 {
			idx++
		}
		if idx < 0 {
			idx = 0
		}
		res = append(res, nums[idx])
		total, k = t, k-t*idx
		nums = append(nums[0:idx], nums[idx+1:]...)
	}
	return string(res)
}
