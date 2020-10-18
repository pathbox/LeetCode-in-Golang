package LeetCode621

func leastInterval(tasks []byte, n int) int {
	charSlice := [26]int{}
	max := 0
	count := 0
	for i := 0; i < len(tasks); i++ {
		cur := tasks[i] - 'A'
		charSlice[cur]++
		if max < charSlice[cur] {
			max = charSlice[cur]
			count = 1
		} else if charSlice[cur] == max {
			count++
		}
	}
	if n == 0 || (max-1)*(n+1)+count < len(tasks) {
		return len(tasks)
	}
	return (max-1)*(n+1) + count
}
