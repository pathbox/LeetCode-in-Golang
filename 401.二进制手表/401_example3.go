package LeetCode401

import "fmt"

/*
按照回溯法的模板进行解题。
这里用字符串不是很便利，用位运算比较便利
func tree(选择,路径){
   结束条件
   遍历分叉
     if 满足剪枝条件
        continue
     进入节点前干啥
     递归节点
     遍历节点后干啥
}

*/
var watches []string

func readBinaryWatch(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}
	watches = make([]string, 0)
	watchHelper(num, make([]int, 10), 0)
	return watches
}

func watchHelper(rec int, bits []int, start int) {
	end := 10 - rec
	for start <= end {
		bits[start] = 1
		if rec > 1 {
			watchHelper(rec-1, bits, start+1)
		} else { // 当前取了n个数了开始操作这n个数看能组合得到什么结果
			hour := 0
			var i uint32 = 0
			for ; i < 4; i++ {
				hour += bits[i] << (3 - i)
			}
			minute := 0
			for ; i < 10; i++ {
				minute += bits[i] << (9 - i)
			}
			if hour < 12 && minute < 60 {
				watches = append(watches, fmt.Sprintf("%d:%0.2d", hour, minute))
			}
		}
		bits[start] = 0
		start++
	}
}
