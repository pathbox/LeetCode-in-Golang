package LeetCode029

/*
判断商的正负性；
计算被除数可以减去多少个除数；
若商大于2^31 -1，则返回2^31-1
*/

func divide(dividend int, divisor int) int {
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	k := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	var ans int
	for dividend >= divisor {
		var cnt int
		var tmp int
		for dividend >= divisor<<tmp {
			cnt = tmp
			tmp++
		}
		dividend -= divisor << cnt
		ans += 1 << cnt
	}
	if k {
		return ans
	}
	return -ans
}
