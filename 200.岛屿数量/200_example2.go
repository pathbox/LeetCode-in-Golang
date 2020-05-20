package LeetCode200

var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, 1, -1}
var row, col int

func numIslands(grid [][]byte) int {
	row = len(grid)
	if row == 0 {
		return 0
	}

	col = len(grid[0])
	count := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				BFS(grid, i, j) // 以当前(i,j)节点继续进行BFS遍历直到达到条件结束，说明以当前(i,j)节点遍历形成的岛屿遍历结束，找到一个岛屿结构
				count++
			}
		}
	}
	return count
}

func BFS(grid [][]byte, i, j int) {
	queue := make([]int, 0)
	queue = append(queue, i, j)
	grid[i][j] = '0'
	for len(queue) != 0 {
		i, j := queue[0], queue[1]
		queue = queue[2:]
		for m := 0; m < 4; m++ {
			tmpI := i + dx[m] // 遍历四个方向
			tmpJ := j + dy[m]

			if 0 <= tmpI && tmpI < row && tmpJ >= 0 && tmpJ < col && grid[tmpI][tmpJ] == '1' {
				grid[tmpI][tmpJ] = '0'
				queue = append(queue, tmpI, tmpJ)
			}
		}
	}
}
