package LeetCode306

import "strconv"

func isAdditiveNumber(num string) bool {
	n := len(num)
	if n < 3 {
		return false
	}
	for i := 1; i < n; i++ {
		if i > 1 && num[0] == '0' { // 排除前导 0
			break
		}
		if lst, _ := strconv.Atoi(num[:i]); dfs(num, lst, i) { // 不是很优雅的递归
			return true
		}
	}
	return false
}

func dfs(num string, lst int, idx int) bool {
	n := len(num)
	isAdditive := false
	for nxtStart := idx + 1; nxtStart < n; nxtStart++ {
		if nxtStart > idx+1 && num[idx] == '0' { // 排除前导 0
			break
		}
		cur, _ := strconv.Atoi(num[idx:nxtStart]) // 第二个数
		sum := strconv.Itoa(lst + cur)            // 第一个数和第二个数的和
		dur := len(sum)                           // 第三个数的长度
		nxtEnd := nxtStart + dur                  // 第三个数的右区间
		if nxtEnd > n {                           // 如果右区间越界，结束循环
			break
		}
		if sum == num[nxtStart:nxtEnd] { // 如果和相等
			if nxtEnd == n { // 并且刚好到数组的最后一位，则为累加数
				return true
			}
			isAdditive = dfs(num, cur, nxtStart) // 没到结尾就继续递归
		}
	}
	return isAdditive
}
