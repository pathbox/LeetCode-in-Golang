package LeetCode054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	dir := 1 // 初始向右
	row, col := 0, 0
	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0 // 初始四个边界,每次一个边的操作，四个边界不断的往内收缩
	var res []int
	for top <= bottom && left <= right { // 循环条件
		res = append(res, matrix[row][col])
		switch dir {
		case 1:
			if col == right {
				dir = 2
				top++
				row++
				continue
			}
			col++
		case 2:
			if row == bottom {
				dir = 3
				right--
				col--
				continue
			}
			row++
		case 3:
			if col == left {
				dir = 4
				bottom--
				row--
				continue
			}
			col--
		case 4:
			if row == top {
				dir = 1
				left++
				col++
				continue
			}
			row--
		}
	}
	return res
}

/*
dir
1、向右
2、向下
3、向左
4、向上
向右运动结束，就会自动向下运动，接着向左，向上，这个很好理解

2、每次运动完后，边界都会发生变化


1、向右运动结束后，上边界下移，即top++
2、向下运动结束后，右边界左移，即right--
3、向左运动结束后，下边界上移，即bottom--
4、向上运动结束后，左边界右移，即left++
3、 每次运动结束后，开始新方向运动时，对应的开始行列也会发生变化


1、向右运动结束后，开始行下移，即row++
2、向下运动结束后，开始列左移，即col--
3、向左运动结束后，开始行上移，即row--
4、向上运动结束后，开始列右移，即col++

*/
// https://leetcode-cn.com/problems/spiral-matrix/solution/golang-qian-xian-yi-dong-qie-shuang-bai-by-bloodbo/
