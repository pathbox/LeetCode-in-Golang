package LeetCode557

import "fmt"

func reverseWords(s string) string {
	array := []byte(s)
	fmt.Println(array)
	start := 0
	for i := 0; i < len(array); i++ {
		if array[i] == ' ' {
			reverse(array, start, i-1)
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
