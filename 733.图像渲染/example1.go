package LeetCode733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	origColor := image[sr][sc]
	fill(image, sr, sc, origColor, newColor)
	return image
}

func fill(image [][]int, x, y, origColor, newColor int) {
	if !inArea(image, x, y) {
		return
	}

	if image[x][y] != origColor { // 其他颜色不会修改
		return // 碰壁：遇到其他颜⾊，超出 origColor 区域
	}
	if image[x][y] == -1 { // 已经探索过，返回
		return // // 已探索过的 origColor 区域
	}

	// 已探索过的 origColor 区域
	image[x][y] = -1 // 设置为-1 表示已经探索过，之后再第29行，将设置为newColor，为什么可以呢？因为29行是回溯结束，递归出栈的时候这一层最后执行的，由于在上面的3个过滤条件都通过了，所以这个x，y是要给上newColor的

	fill(image, x+1, y, origColor, newColor)
	fill(image, x-1, y, origColor, newColor)
	fill(image, x, y+1, origColor, newColor)
	fill(image, x, y-1, origColor, newColor)
	// unchoose：将标记替换为 newColor
	image[x][y] = newColor
}

// 是否在区域中
func inArea(image [][]int, x, y int) bool {
	return x >= 0 && x < len(image) && y >= 0 && y < len(image[0])
}
