package LeetCode542

func updateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])

	queue := make([][]int, 0)
	for i := 0; i < n; i++ { // 把0全部存进队列，后面从队列中取出来，判断每个访问过的节点的上下左右，直到所有的节点都被访问过为止。
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				point := []int{i, j}
				queue = append(queue, point) // 0是所有初始出发点
			} else {
				matrix[i][j] = -1
			}
		}
	}
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // 上下左右遍历的好方法
	for len(queue) > 0 {                                   // 这里就是 BFS 模板操作了
		point := queue[0]
		queue = queue[1:]
		for _, v := range direction {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y}) // 计算过的x，y会放入队列，以x,y为起始点继续向外找周围是-1(1)的情况的值。因为matrix[x][y]的值为x,y到附件最近的0点的距离,所以如果以x,y点上下左右移动一位，得到的matrix[x'][y']= matrix[x][y]+1
			}
		}
	}
	return matrix
}

func dfs(matrix [][]int, cm *[][]int, i, j int) {
	if matrix[i][j] == 0 {
		cm[i][j] = 0
	}

}
