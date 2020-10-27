package LeetCode886

// https://leetcode-cn.com/problems/possible-bipartition/solution/golangbfshe-dfsran-se-by-bloodborne/

func possibleBipartition(N int, dislikes [][]int) bool {
	// 维护每个人讨厌的人物
	graph := make([][]int, N+1) // 邻接矩阵
	for _, v := range dislikes {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	color := make([]int, N+1)
	for i := 1; i <= N; i++ {
		// 如果已经染色了，说明其所在的讨厌圈子已经处理了，不用再看了
		if color[i] != 0 {
			continue
		}

		color[i] = 1 // 新的讨厌圈子第一位染色1就好了
		queue := []int{i}
		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]
			for _, v := range graph[head] {
				// 两边都不容，无处容身！
				if color[v] == color[head] { // 1.head 和其所讨厌的人的颜色不能一样，不能在一组
					return false
				}
				if color[v] != 0 { // 2.表示已经染色
					continue
				}
				if color[v] == 0 {
					color[v] = -color[head] // 3.v的颜色和head的颜色要是相反的
				}
				queue = append(queue, v) // 入队列，需要等会遍历v的讨厌的人的情况
			}
		}
	}
	return true
}
