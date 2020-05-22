package LeetCode695

// nice: https://leetcode-cn.com/problems/max-area-of-island/solution/fang-ge-lei-dfs-de-jian-dan-fang-fa-cjava-by-nette/

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}
	var max int
	h := len(grid)
	w := len(grid[0])
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 1 {
				temp := dfs(grid, i, j)
				if temp > max {
					max = temp
				}
			}
		}
	}
	return max
}

func dfs(grid [][]int, x, y int) int {
	h := len(grid)
	w := len(grid[0])
	var res int

	if grid[x][y] == 1 {
		grid[x][y] = 0 // 设为2也可以，只要不保持为1
		// return 1 + dfs(grid, x-1, y) + dfs(grid, x+1, y) + dfs(grid, x, y-1) + dfs(grid, x, y+1) // 四个方向遍历累加

		if x-1 >= 0 && grid[x-1][y] == 1 { //上节点为1
			res += dfs(grid, x-1, y)
		}
		if x+1 < h && grid[x+1][y] == 1 { //上节点为1
			res += dfs(grid, x+1, y)
		}
		if y-1 >= 0 && grid[x][y-1] == 1 { //上节点为1
			res += dfs(grid, x, y-1)
		}
		if y+1 < w && grid[x][y+1] == 1 { //上节点为1
			res += dfs(grid, x, y+1)
		}
		return res + 1
	}
	return res
}
