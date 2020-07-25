package LeetCode186

import "fmt"

// 思路： 两次反转
//1，将整个字符串进行反转
//2，再将每个单词进行反转
//3，注意边界条件的处理：最后一个字符的处理
func reverseWords(s []byte) {
	if len(s) == 0 {
		return
	}

	reverseString(s)
	fmt.Println(s)
	preIndex := 0
	for index := 0; index < len(s); index++ {
		if s[index] == ' ' {
			reverseString(s[preIndex:index])
			preIndex = index + 1
		}
		//这里需要判断末尾的边界 对末尾的处理
		if index == len(s)-1 {
			reverseString(s[preIndex : index+1])
		}
	}
	fmt.Println(s)
}

func reverseString(s []byte) {
	front, back := 0, len(s)-1
	for front < back {
		s[front], s[back] = s[back], s[front]
		front++
		back--
	}
}
