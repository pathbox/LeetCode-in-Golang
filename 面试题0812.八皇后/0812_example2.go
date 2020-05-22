package LeetCode0812

func solveNQueens(n int) [][]string {
	nums := make([][]rune, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			nums[i] = append(nums[i], '.')
		}
	}
	res := make([][]string, 0)
	helpSolveNQueens(&res, nums, 0)
	return res
}

func helpSolveNQueens(res *[][]string, nums [][]rune, curRow int) {
	n := len(nums)
	if curRow == n {
		tmp := make([]string, 0)
		for i := 0; i < len(nums); i++ {
			tmp = append(tmp, string(nums[i]))
		}
		*res = append(*res, tmp)
		return
	}
	//当前行 放置 Q 的所有情况的遍历（也就是对当前行的 所有列进行遍历，寻找所有可以放置 Q 的 当前行的 列，并向下传递），遍历完之后 将 每一种情况 传递到下一行 继续递归遍历
	for col := 0; col < n; col++ {

		flag := true
		//只用关心curRow前面的 所有的行里面 同列 对角线 有没有Q，因为curRow以及后面的行里面都是 . . . . .没有Q，还没有遍历到后面的行
		for row := 0; row < curRow; row++ {
			//如果前面的所有行中 相同列 有Q则跳过这一列的 所有行遍历，到下一列继续寻找成立的 条件
			if nums[row][col] == 'Q' {
				flag = false
				break
			}
			if col+curRow-row < n && nums[row][col+curRow-row] == 'Q' {
				flag = false
				break
			}
			if col-(curRow-row) >= 0 && nums[row][col-(curRow-row)] == 'Q' {
				flag = false
				break
			}
		}
		//这一列的前面 所有行中 存在 同列或者前后对角线 有Q，也就是当前 列 不可用 即 点 curRow,col不可用，试试下一个col 是否可用
		if !flag {
			continue
		}
		//如果遍历了前面所有行 flag都没有被置成 false 说明当前行的 这一列 的所有行，所有前后对角线，都没有Q，也就是当前位置curRow col 符合条件
		nums[curRow][col] = 'Q'
		helpSolveNQueens(res, nums, curRow+1)
		//还原 继续看 当前行的其他的列是不是也可以找到符合条件的 可以放置 Q的  点
		nums[curRow][col] = '.'
	}
}
