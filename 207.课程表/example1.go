package LeetCode207

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	result := 0
	queue := make([]int, 0)

	//建立邻接矩阵，入度表
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		inDegree[v[0]]++
	}

	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
			//设置为无效
			inDegree[k] = -1
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		result++
		queue = queue[1:]
		for _, v := range edges[node] {
			if inDegree[v] == 1 {
				queue = append(queue, v)
				//设置为无效
				inDegree[v] = -1
			} else {
				inDegree[v]--
			}
		}

	}
	return result == numCourses
}
