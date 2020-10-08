package LeetCode042

func trap(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				//设置左边最高柱子
				leftMax = height[left]
			} else {
				//右边必定有柱子挡水，所以，遇到所有值小于等于leftMax的，全部加入水池
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				//设置右边最高柱子
				rightMax = height[right]
			} else {
				//左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}
