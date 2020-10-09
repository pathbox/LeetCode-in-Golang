package LeetCode134

// 反证思想法 要么无解 要么肯定会有一个解，不需要遍历求出所有情况，只要找到一个情况即可
func canCompleteCircuit(gas []int, cost []int) int {
	// 数组长度 / 油总量 - 消耗量重量 / 当前汽车油箱油量 / 最新的起点
	l, sum, tank, startIndex := len(gas), 0, 0, 0
	for i := 0; i < l; i++ {
		sum += gas[i] - cost[i]
		tank += gas[i] - cost[i]
		// 汽车油还够用，继续走着
		if tank >= 0 {
			continue
		}
		// 当前汽车油量为负数了，说明这个加油站不能作为起点 排除法
		startIndex = i + 1 // 改立下一个吧
		tank = 0           // 清空当前汽车油箱
	}
	// 油总量 - 消耗量重量 为负，说明油根本不够消耗的
	if sum < 0 {
		return -1
	}
	// 要不然一定有解
	return startIndex
}
