package LeetCode084

/*
单调栈
只有当前栈为空或者当前高度大于栈顶值的高度，才能入栈
如果当前值等于栈顶值的高度，不做任何处理
如果当前值小于栈顶值的高度，那么将当前值弹出站，计算弹出的值对应的最大面积，更新最大面积，如果并记录弹出的开始位置；如果栈顶的值等于当前值，不做处理；如果当前值大于栈顶值或者栈顶为空，将当前值push到栈里面，并且记录起始位置为最后弹出的值的起始位置
遍历完毕，如果栈不为空，将栈元素弹出，并以此记录每个元素对应的最大面积

对于每个高度来说，往左右两个方向寻找第一个比它矮的柱子，决定矩形的宽。
    // 所以用高度升序栈保存索引，遇到减小的，表示找到右边界。计算这中间每个柱子的最大矩形，计算一个即出栈，这样方便每次从栈顶取索引。
    // 时间复杂度：O(n)， n个数字每个会被压栈弹栈各一次。
    // 空间复杂度：O(n)。用来存放栈中元素
*/

func largestRectangleArea(heights []int) int {
	// 首尾添加负数高度，这样原本的第一个高度能形成升序，原本的最后一个高度也能得到处理
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)
	size := len(heights)
	// 递增栈
	s := make([]int, 1, size) // 存的是索引位置

	res := 0
	i := 1
	for i < len(heights) {
		if heights[s[len(s)-1]] < heights[i] {
			s = append(s, i)
			i++
			continue
		}
		// s[len(s)-2]是矩形的左边界
		res = max(res, heights[s[len(s)-1]]*(i-s[len(s)-2]-1))
		s = s[:len(s)-1] // 出栈
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
