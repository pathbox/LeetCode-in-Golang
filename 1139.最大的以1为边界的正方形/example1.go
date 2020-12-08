package LeetCode1139

import "math"

//  https://leetcode-cn.com/problems/largest-1-bordered-square/solution/971100-by-stay-pain/
func largest1BorderedSquare(grid [][]int) int {
	//如果是一维数组
	if len(grid) == 1 {
		for i := 0; i < len(grid); i++ {
			for _, value := range grid[i] {
				if value == 1 {
					return 1
				}
			}
		}
	}
	//	如果是多维数组
	length := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			temp := maxSquare(grid, i, j, length)
			if temp > length {
				length = temp
			}
		}
	}
	return length * length
}

func maxSquare(grid [][]int, row, col, length int) int {
	top, bottom, left, right := 0, 0, 0, 0
	// 先确定上、左边界
	//	上边界
	for i := col; i < len(grid[0]); i++ {
		if grid[row][i] == 1 {
			top++
		} else {
			break
		}
	}
	//	左边界
	for i := row; i < len(grid); i++ {
		if grid[i][col] == 1 {
			left++
		} else {
			break
		}
	}

	// 根据上左边界来确定边长，以此为标准来寻找另外两天边
	//	目前可能的最大的正方形的边长为min(left, top)
	tempLength := int(math.Min(float64(top), float64(left)))
	if length > tempLength {
		return length
	}

	for tempLength != bottom || tempLength != right {
		//下边界
		for i := col; i <= col+tempLength-1; i++ {
			if grid[row+tempLength-1][i] == 1 {
				bottom++
			} else {
				break
			}
		}
		//右边界
		for i := row; i <= row+tempLength-1; i++ {
			if grid[i][col+tempLength-1] == 1 {
				right++
			} else {
				break
			}
		}
		// 无法围成正方形， tempLength缩小1单位继续寻找
		if right != bottom || right != tempLength {
			bottom = 0
			right = 0
			tempLength--
			continue
		}
		if tempLength == 0 {
			break
		}
	}
	return tempLength
}
