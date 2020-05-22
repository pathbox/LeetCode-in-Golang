package LeetCode0812

func solveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}
	col := make([]bool, n)
	dig1, dig2 := make([]bool, n<<1-1), make([]bool, n<<1-1)
	var res [][]string
	var tmp [][]byte
	tmp = make([][]byte, n)
	for i := range tmp {
		tmp[i] = make([]byte, n)
		for j := range tmp[i] {
			tmp[i][j] = '.'
		}
	}
	var dfs func(h int)
	dfs = func(h int) {
		if h == n {
			t := make([]string, n)
			for i := range t {
				t[i] = string(tmp[i])
			}
			res = append(res, t)
		} else {
			for i := 0; i < n; i++ {
				if !col[i] && !dig1[i-h+n-1] && !dig2[i+h] {
					col[i], dig1[i-h+n-1], dig2[i+h] = true, true, true
					tmp[h][i] = 'Q'
					dfs(h + 1)
					tmp[h][i] = '.'
					col[i], dig1[i-h+n-1], dig2[i+h] = false, false, false
				}
			}
		}
	}
	dfs(0)
	return res
}
