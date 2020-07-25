package LeetCode672

// https://leetcode-cn.com/problems/bulb-switcher-ii/solution/ying-zhao-chu-lai-de-gui-lu-shuang-bai-tong-guo-by/

func flipLights(n, m int) int {
	if n == 0 {
		return 0
	}
	if n >= 3 {
		return min(8, 1+m*3)
	}
	if n == 1 {
		return min(2, 1+m)
	}
	if n == 2 {
		return min(4, 1+m*2)
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
