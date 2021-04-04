package LeetCode456

// 栈中的元素是max  出栈的min其实是mid， 再往左遍历的是min，结果就是 min max mid的排序 满足132模式
func find132pattern(nums []int) bool {
	stack := make([]int, 0)
    min := -1 << 63

		for k := len(nums)- 1; k >= 0; k-- {
			n := nums[k]
			if n < min {
				return true
			}

			for len(stack) > 0 && n > stack[len(stack)-1] {
				i := len(stack) -1
				min = max(min, stack[i])
				stack = stack[:i]
			}

			stack = append(stack, n)
		}
		return false
}
