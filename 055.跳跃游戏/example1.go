package LeetCode055

// 贪心 遍历
func canJump(nums []int) bool {
	max, ln := 0, len(nums)-1
	for i := 0; i <= ln; i++ {
		if i <= max { // 当前下标在能跳到的max内
			tmp := i + nums[i] // 当前位置能跳到多远
			if tmp >= ln {     // 如果能跳过最后一个元素则说明可以跳过
				return true
			}
			if tmp > max {
				max = tmp // 更新最远跳到位置
			}
		} else { // 当前位置跳不到,提前结束, i > max 说明前面怎么跳都是跳不到当前位置的，索引返回false
			return false
		}
	}
	return false
}
