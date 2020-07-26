package LeetCode256

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	// 3 表示3种颜色的选择
	var prev [3]int          // 上一个为止的最小cost和
	var curr [3]int          // 当前的cost
	for i := 0; i < 3; i++ { // 初始化第一个房子的三种颜色的cost
		prev[i] = costs[0][i]
	}

	for i := 1; i < len(costs); i++ {
		curr[0] = costs[i][0] + min(prev[1], prev[2]) // 当前选择红色的cost=红色cost+上一个为止的最小cost和
		curr[1] = costs[i][1] + min(prev[0], prev[1])
		curr[2] = costs[i][2] + min(prev[0], prev[1])
		//...同样的可以扩充到多种颜色,然后将min方法改成支持数组中找到最小值的方式，初始化的时候将所有颜色初始化好
		prev = curr
	}
	return min(min(prev[0], prev[1]), prev[2]) // 返回prev中最小的值就是最小的cost

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
