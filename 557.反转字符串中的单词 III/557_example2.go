package LeetCode557

import "fmt"

func reverseWords(s string) string {
	array := []byte(s)
	fmt.Println(array)
	start := 0
	for i := 0; i < len(array); i++ {
		if array[i] == ' ' { // 此时 array[:i]是一个单词
			reverse(array, start, i-1) // 进行该单词的头尾交换
			start = i + 1
		}
	}
	reverse(array, start, len(array)-1)
	return string(array)
}

func reverse(array []byte, start, end int) {
	for start < end {
		array[start], array[end] = array[end], array[start]
		start++
		end--
	}
}
