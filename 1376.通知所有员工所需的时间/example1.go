package LeetCode1376

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	managerMap := make(map[int][]int)

	// 所有上下级的关系
	for i, v := range manager {
		if _, ok := managerMap[v]; ok {
			managerMap[v] = append(managerMap[v], i)
		} else {
			managerMap[v] = []int{i}
		}
	}

	max := 0
	cost := 0
	var search func(supId int)
	search = func(supId int) {
		subs, ok := managerMap[supId]
		if !ok {
			if cost > max {
				max = cost
			}
			return
		}

		cost += informTime[supId]
		for _, v := range subs {
			// v是下属headId
			search(v)
		}
		// 回溯
		cost -= informTime[supId]
	}

	search(headID)
	return max
}
