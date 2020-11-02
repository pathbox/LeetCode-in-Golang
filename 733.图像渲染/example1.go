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

	if image[x][y] != origColor {
		return // 碰壁：遇到其他颜⾊，超出 origColor 区域
	}
	if image[x][y] == -1 {
		return // // 已探索过的 origColor 区域
	}

	// 已探索过的 origColor 区域
	image[x][y] = -1

	fill(image, x+1, y, origColor, newColor)
	fill(image, x-1, y, origColor, newColor)
	fill(image, x, y+1, origColor, newColor)
	fill(image, x, y-1, origColor, newColor)
	// unchoose：将标记替换为 newColor
	image[x][y] = newColor
}

func inArea(image [][]int, x, y int) bool {
	return x >= 0 && x < len(image) && y >= 0 && y < len(image[0])
}
