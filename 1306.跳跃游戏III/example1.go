package LeetCode1306

func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i < 0 || i > len(arr)-1 || visited[i] { // 越界 以及访问过都返回false
			return false
		}
		visited[i] = true
		return arr[i] == 0 || dfs(i+arr[i]) || dfs(i-arr[i])
	}
	return dfs(start)
}
