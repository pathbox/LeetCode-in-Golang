package main

import "fmt"

func reverseWords(s string) string {
	array := []byte(s)
	start := 0
	for i := 0; i < len(array); i++ {
		if array[i] == ' ' { // 此时 array[:i]是一个单词
			reverse(array, start, i-1) // 进行该单词的头尾交换,此时的i-1是空格字符的前面一个索引，即一个单词的长度
			start = i + 1
		}
	}
	// 上面只处理到' '之前的字符串，还有最后一个单词没有处理
	fmt.Println(start, string(array))
	reverse(array, start, len(array)-1)
	fmt.Println(start, string(array))
	return string(array)
}

func reverse(array []byte, start, end int) {
	for start < end {
		array[start], array[end] = array[end], array[start]
		start++
		end--
	}
}

func main() {
	reverseWords("Let's take LeetCode contest")
}
