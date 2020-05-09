package LeetCode401

import "strconv"

func readBinaryWatch(num int) []string {
	hours := [4]int{1, 2, 4, 8}
	minutes := [6]int{1, 2, 4, 8, 16, 32}
	res := make([]string, 0)
	//接受转换后的字符串
	s := ""
	//计算小时
	shour := 0
	//计算分钟
	sminute := 0

	if num == 0 {
		s = "0:00"
		res = append(res, s)
		return res
	}

	Run(hours, minutes, s, &res, shour, sminute, 0, 0, num)
	return res
}

func Run(hours [4]int, minutes [6]int, s string, res *[]string, shour, sminute, h, m, num int) {
	if num == 0 {
		s += strconv.Itoa(shour)
		s += ":"
		if sminute <= 9 {
			s += "0"
		}
		s += strconv.Itoa(sminute)
		*res = append(*res, s)
		return
	}
	// 分钟
	for i := m; i < 6; i++ {
		if sminute+minutes[i] <= 59 {
			//再回溯回来的时候m之前不再进行选择
			//因为已经选择并判断过
			m++
			Run(hours, minutes, s, res, shour, sminute+minutes[i], h, m, num-1)
		}
	}

	//小时
	for i := h; i < 4; i++ {
		if shour+hours[i] <= 11 {
			h++
			Run(hours, minutes, s, res, shour+hours[i], sminute, h, m, num-1)
		}
	}
}
