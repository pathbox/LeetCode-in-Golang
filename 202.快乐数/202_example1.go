// package LeetCode202
package main

import "fmt"

func main() {
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	m := make(map[int]struct{})
	for n != 1 {
		a := getArray(n)
		n = sum(a)
		_, ok := m[n]
		if ok {
			return false
		} else {
			m[n] = struct{}{}
		}
	}
	return true
}

func getArray(n int) []int {
	a := make([]int, 0)
	for n != 0 {
		i := n % 10
		a = append(a, i)
		n = n / 10
	}
	return a
}

func sum(a []int) int {
	r := 0
	for _, i := range a {
		r = r + i*i
	}
	return r
}

// 如果出现了平方元素组成的循环，则说明不是快乐数
