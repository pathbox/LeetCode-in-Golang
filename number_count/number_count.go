/*
一个数组 array [1,3,5,9,8,7,6,55,23,55]
一个数 num， 计算该num在array中出现的次数
*/

package main

import (
	"fmt"

	"../mysort"
)

func main() {
	array := []int{1, 3, 5, 9, 8, 7, 6, 55, 23, 55}
	num := 55

	low := 0
	hight := len(array) - 1

	mysort.QuickSort(low, hight, array)

	fmt.Println("array: ", array)

	result := numCount(array, num)
	fmt.Printf("%d count in array: %d", num, result)
}

func numCount(array []int, num int) int {
	lo, hi := 0, len(array)-1
	c := 0
	for lo <= hi {
		m := (lo + hi) >> 1
		if array[m] < num {
			lo = m + 1
		} else if array[m] > num {
			hi = m - 1
		} else {
			c++ // 找到要找的数了， c++ 一次
			mid := m
			//以找到的m index为坐标，往前走，因为array已经排序了，所以如果array[index]值不相等num，则中断循环，相等继续往前走直到（左）头
			m = m - 1
			for m >= 0 && array[m] == num {
				m--
				c++
			}
			fmt.Println("mid", mid)
			//以找到的m index为坐标，往后走，因为array已经排序了，所以如果array[index]值不相等num，则中断循环，相等继续往后走直到（右）头
			mid = mid + 1
			for mid <= len(array)-1 && array[mid] == num {
				mid++
				c++
			}
			fmt.Println("count c ", c)
			return c
		}
	}
	return c
}
