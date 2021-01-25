package Offer

// https://www.bilibili.com/video/BV15D4y1X7Tt?p=16
// 一次删除两个不同的数
func halfMajor(arr []int) int {
	cand := 0
	hp := 0
	// 哪一个数剩下来
	for i := 0; i < len(arr); i++ {
		if hp == 0 {
			cand = arr[i]
			hp = 1
		} else if arr[i] == cand {
			hp++
		} else {
			hp--
		}
	}

	if hp == 0 {
		return -1
	}

	// 验证这个数的数量是否是大于N/2个数
	hp = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == cand {
			hp++
		}
	}
	if hp > len(arr)/2 {
		return cand
	}
	return -1
}
