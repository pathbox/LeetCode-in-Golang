package LeetCode564

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	// 首先判断是否为10的整数倍，此情况比较特殊
	if v, ok := tenTimes(n); ok { // 如果是，最后返回的是9(999...)这样的数
		return v
	}
	// 获取镜像, mirror1、mirror2为两种情况
	mirror1 := getMirror(n)
	var mirror2 string
	switch compare(n, mirror1) { // 会在中间的位置进行+1 或 -1，之后再以此值进行镜像化，得到一个回文数
	case -1:
		// 镜像大,-1
		mirror2 = getMirror(add(n, -1))
	case 0:
		// 相等,+1,-1
		mirror1 = getMirror(add(n, -1))
		mirror2 = getMirror(add(n, 1))
	case 1:
		// 镜像小,+1
		mirror2 = getMirror(add(n, 1))
	}
	// 比较mirror1和mirror2与n差的绝对值，返回差值较小的；如果相等，则返回mirror1和mirror2中较小的
	m1, m2 := minor(mirror1, n), minor(mirror2, n)
	if m1 < m2 {
		return mirror1
	} else if m1 == m2 {
		if compare(mirror1, mirror2) < 0 {
			return mirror1
		}
		return mirror2
	}
	return mirror2
}

func tenTimes(s string) (string, bool) {
	sInt, _ := strconv.Atoi(s)
	if (sInt-1)%10 == 0 {
		sInt = sInt - 1
	}
	i := sInt
	for ; i >= 10; i = i / 10 {
		if i%10 != 0 {
			return "", false
		}
	}
	if i != 1 {
		return "", false
	}
	return strconv.Itoa(sInt - 1), true
}

func add(s string, number int) string {
	sByte := []byte(s)
	i, j := 0, len(sByte)-1
	for i <= j {
		i, j = i+1, j-1
	}
	// i 此时是中间索引值
	sInt, _ := strconv.Atoi(string(sByte[0:i]))
	sAdd := strconv.Itoa(sInt + number)
	if i == len(sByte) {
		return sAdd
	}
	return sAdd + string(s[i:]) // 加完后再将后半部分拼接
}

func minor(l, r string) int {
	lInt, _ := strconv.Atoi(l)
	rInt, _ := strconv.Atoi(r)
	return int(math.Abs(float64(lInt - rInt)))
}

func compare(l string, r string) int {
	lInt, _ := strconv.Atoi(l)
	rInt, _ := strconv.Atoi(r)
	if lInt > rInt {
		return 1
	} else if lInt == rInt {
		return 0
	}
	return -1
}
func getMirror(s string) string {
	sByte := []byte(s)
	for i, j := 0, len(sByte)-1; i <= j; i, j = i+1, j-1 { // 以前半部分构造镜像
		sByte[j] = sByte[i]
	}
	return string(sByte)
}

/*
"1121211"
得到
"1120211"
*/
