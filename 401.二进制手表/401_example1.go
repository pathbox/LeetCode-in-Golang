package LeetCode401

import "fmt"

func readBinaryWatch(num int) []string {
	ret := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j <= 59; j++ {
			if bitNums(i)+bitNums(j) == num {
				ret = append(ret, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return ret
}

func bitNums(i int) int {
	num := 0
	for i > 0 {
		if i%2 == 1 {
			num++
		}
		i = i >> 1
	}
	return num
}
