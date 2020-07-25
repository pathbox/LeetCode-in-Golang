package LeetCode593

import "sort"

/*
4边相等
对角线是边长的平方的2倍
把六条线长度平方算出来，如果有4条是相等的，且剩余两条是这四条的两倍，
那一定是一个正方形
https://leetcode-cn.com/problems/valid-square/solution/you-xiao-de-zheng-fang-xing-by-leetcode/
*/
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	data := [][]int{p1, p2, p3, p4}
	sort.Slice(data, func(i, j int) bool { // 我们将这 4 个点按照 x 轴坐标第一关键字，y 轴坐标第二关键字进行升序排序
		if data[i][0] < data[j][0] {
			return true
		}
		if data[i][0] > data[j][0] {
			return false
		}

		return data[i][1] < data[j][1]
	})

	length := sqrLen(data[0], data[1])
	if length == 0 {
		return false
	}
	if length != sqrLen(data[0], data[2]) {
		return false
	}

	if length != sqrLen(data[3], data[1]) {
		return false
	}

	if length != sqrLen(data[3], data[2]) {
		return false
	}

	if 2*length != sqrLen(data[1], data[2]) {
		return false
	}

	return true
}

func sqrLen(p1 []int, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

/*

class Solution:
    def dis(self, p1: List[int], p2: List[int]):
        return (p1[0] - p2[0])**2 + (p1[1] - p2[1])**2

    def validSquare(self, p1: List[int], p2: List[int], p3: List[int], p4: List[int]) -> bool:
        l = []
        l.append(self.dis(p1, p2))
        l.append(self.dis(p1, p3))
        l.append(self.dis(p1, p4))
        l.append(self.dis(p2, p3))
        l.append(self.dis(p2, p4))
        l.append(self.dis(p3, p4))
        l.sort()
        return l[0] != 0 and l[1] == l[0] and l[2] == l[1] and l[3] == l[2] and l[4] == l[3]*2 and l[5] == l[4]

*/
