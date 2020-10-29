package LeetCode461

import "fmt"

func hammingDistance(x int, y int) int {
	C := x ^ y
	count := 0

	s := fmt.Sprintf("%b", C)
	for i := range s {
		if s[i] == 49 {
			count++
		}
	}

	return count
}

func hammingDistance(x int, y int) int {
	i := x ^ y   //x和y异或,得到一个新的数(异或相同为0,不同为1,此时咱们需要统计1的个数)
	count := 0   //定义数量的初始值为0
	for i != 0 { //只要i不为0,那就继续循环
		if (i & 1) == 1 { //如果i和1相与,值为1的话就count++
			count++
		}
		i = i >> 1 //i右移一位
	}
	return count

}

func hammingDistance(x int, y int) int {
	C := x ^ y
	count := 0

	for C != 0 {
		C = C & (C - 1)
		count++
	}

	return count
}

/*
n&(n-1)这个操作是算法中常⻅的，作⽤是消除数字 n 的⼆进制表⽰中的最后⼀个 1。
*/
