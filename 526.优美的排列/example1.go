package LeetCode526

var cnt [][]int // 可以在外面定义全局变量 在方法中进行初始化 则不会报错
func countArrangement(N int) int {
	cnt = make([][]int, 0)
	used := make([]int, N)
	tool := make([]int, N)
	for i := 0; i < N; i++ {
		tool[i] = i + 1
	}
	cur := make([]int, 0)
	backtrace(used, cur, N)
	return len(cnt)
}

func backtrace(used, cur []int, N int) {
	if len(cur) == N {
		tmp := make([]int, N)
		copy(tmp, cur)
		cnt = append(cnt, tmp)
		return
	}
	for i := 1; i <= N; i++ {
		if used[i-1] == 0 {
			if i%(len(cur)+1) == 0 || (len(cur)+1)%i == 0 {
				used[i-1] = 1
				cur = append(cur, i)
				backtrace(used, cur, N)
				used[i-1] = 0
				cur = cur[:len(cur)-1]
			}
		}
	}
}
