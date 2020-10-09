package LeetCode134

// https://leetcode-cn.com/problems/gas-station/solution/yi-li-jie-tan-xin-by-dabin-2/

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	// 记录每个加油站加完油+跑到下一个加油站剩余的汽油量
	// 正数：代表这个加油站有利于跑的更远
	// 负数：这个加油站是拖后腿的
	left := make([]int, n, n)
	totalSum := 0
	for i := 0; i < n; i++ {
		left[i] = gas[i] - cost[i]
		totalSum += left[i]
	}

	// 消耗汽油大于总汽油，肯定无解,否则肯定有解
	if totalSum < 0 {
		return -1
	}
	// 从哪个加油站出发
	start := 0
	// 从start位置出发，汽车还剩余的汽油总量
	sum := 0
	for i := 0; i < n; i++ {
		// 当前油箱中汽油为负了，代表从start位置，不足以跑到i位置
		// 得重新寻找加油站。
		// 解会不会在start之前？不会，因为是从前边遍历过来的，如果start前边的位置start'能到
		// 本加油站，start会是start'，而不是当前值
		// 解会不会是start到i位置中的某个加油站？假如存在这样一个加油站start'，start能到start'，
		// 然后到了i，而start还没有切换，说明start到start'时，sum>=0的，有足够的汽油到start'，
		// 如果从start'出发到i，sum只会比当前的汽油更少，可能都不足以到达i。
		// 根据排除法，所以解只能在i后边
		if sum < 0 {
			start = i
			sum = 0
		}
		// 积累油箱中的汽油
		sum += left[i]
	}
	return start
}
