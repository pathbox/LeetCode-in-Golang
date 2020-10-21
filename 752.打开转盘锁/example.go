package LeetCode752

func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	for _, v := range deadends {
		dead[v] = true
	} // 填充dead set
	if dead["0000"] {
		return -1
	} // 直接死锁
	if target == "0000" {
		return 0
	} // 出发即是终点，特殊

	// BFS --------------------------------------------------------------
	queue := make([]string, 0)         // 构造处理字符串队列
	queue = append(queue, "0000")      // 起点
	visited := make(map[string]uint16) // 已访问过的集合。由于总共只有一万个状态点，所以步数不可能需要更多，所以uint16足以表示
	visited["0000"] = 0

	var cur string
	var curSlice []byte
	var nexts [8]string
	var origin byte
	for len(queue) != 0 {
		cur = queue[0]         // 取出当前待处理的锁状态（或者说无向图的节点）
		queue = queue[1:]      // 出队
		curSlice = []byte(cur) // 转为切片

		// 获取当前状态下一步的所有（8个）可能状态
		for i := 0; i < 4; i++ { // 对每一位进行变动。
			origin = curSlice[i] // 备份下原始的字符
			// 正向转动转盘
			curSlice[i] = (curSlice[i]-'0'+1)%10 + '0' // '0'~'9'的字符减去'0' 变为整型，来和1作加减，外边再 + '0'又转为字符
			nexts[2*i] = string(curSlice)
			curSlice[i] = origin // 恢复原始状态
			// 反向转动转盘
			curSlice[i] = (curSlice[i]-'0'+9)%10 + '0' // 如果是-1会出现负数情况，不好处理。循环左移的技巧
			nexts[2*i+1] = string(curSlice)
			curSlice[i] = origin
		}

		// 遍历下一步的所有可能状态
		for _, next := range nexts {
			if _, ok := visited[next]; !ok && !dead[next] { // 没有访问过，也不是dead
				queue = append(queue, next)      // 入队
				visited[next] = visited[cur] + 1 // 步数增加
				if next == target {
					return int(visited[next])
				} // 如果到达目标，就返回最少需要的步数
			}
		}

	}

	return -1
}
