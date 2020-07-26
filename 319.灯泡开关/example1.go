package LeetCode319

import "math"

// https://leetcode-cn.com/problems/bulb-switcher/solution/golang-by-bloodborne/
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
