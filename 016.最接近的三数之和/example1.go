package LeetCode016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if math.Abs(float64(cur-target)) < math.Abs(float64(best-target)) {
			best = cur
		}
	}

	// 三个指针： i j(i+1) k(n-1)  first指针是固定的，只由第一个for循环控制。里面的逻辑控制second指针和third指针的移动
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target { // 完美匹配
				return target
			}
			update(sum)
			if sum > target {
				// 如果和大于 target，移动 c 对应的指针,移动右指针
				k0 := k - 1
				// 移动到下一个不相等的元素
				for j < k0 && nums[k0] == nums[k] { // 继续移动，排除右边重复的数字
					k0--
				}
				k = k0
			} else {
				// 如果和小于 target，移动 b 对应的指针
				j0 := j + 1
				// 移动到下一个不相等的元素
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}
