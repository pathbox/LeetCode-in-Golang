package offer

// https://www.bilibili.com/video/BV15D4y1X7Tt?p=17
func kMajor(arr []int, K int) []int {
	ans := make([]int, 0)
	if K < 2 {
		return ans
	}
	cands := make(map[int]int) // K-1长度
	for i := 0; i < len(arr); i++ {
		if _, ok := cands[arr[i]]; ok {
			cands[arr[i]] = cands[arr[i]] + 1
		} else {
			if len(cands) == K-1 {
				allCandsMinusOne(cands)
			} else {
				cands[arr[i]] = 1
			}
		}
	}

	reals := getReals(arr, cands)
	n := len(arr) / K
	for k, v := range reals {
		if v > n {
			ans = append(ans, k)
		}
	}
	return ans
}

func allCandsMinusOne(cands map[int]int) {
	removeList := make([]int, 0)
	for k, v := range cands {
		if v == 1 {
			removeList = append(removeList, k)
		}
		cands[k] = cands[k] - 1
	}
	// 为1的值 减1后为0，要删掉的
	for _, v := range removeList {
		delete(cands, v)
	}
}

func getReals(arr []int, cands map[int]int) map[int]int {
	reals := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		curNum := arr[i]
		if _, ok := cands[curNum]; ok {
			if _, ok := reals[curNum]; ok {
				reals[curNum] = reals[curNum] + 1
			} else {
				reals[curNum] = 1
			}
		}
	}
	return reals
}
