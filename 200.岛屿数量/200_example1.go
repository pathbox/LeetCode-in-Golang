package LeetCode200

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' { // 1 表示有陆地， 每次代入i，j位置的陆地，以当前这块陆地出发，进行上下左右的探测是否仍然为陆地,探测过的陆地会被设置为0.如果这些陆地是连在一起的，那么实际上还是一块，下一次取 i，j的时候访问过的陆地就不会在访问，它已经和上一次访问的陆地连在一块
				//grid[i][j] == '1' 说明是陆地肯定是岛屿，helper的作用就是把这个岛屿的面积探测完，置为1，避免之后重复探测
				helper(i, j, grid) // 结束一次dfs表示，一块陆地探测完毕，所以结果res++
				res++
			}
		}
	}
	return res
}

func helper(i, j int, grid [][]byte) {
	grid[i][j] = '0' // i，j这块陆地已经访问过,以当前这块陆地出发，进行上下左右的探测是否仍然为陆地

	if i+1 < len(grid) && grid[i+1][j] == '1' {
		helper(i+1, j, grid)
	}

	if i-1 >= 0 && grid[i-1][j] == '1' {
		helper(i-1, j, grid)
	}

	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		helper(i, j+1, grid)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		helper(i, j-1, grid)
	}
}
