package LeetCode470

func rand10() int {
	for {
		r1 := rand7()
		r2 := rand7()
		num := r1 + (r2-1)*7 // 等概率生成[1,49]范围的随机数
		if num <= 40 {       // 拒绝采样，10的倍数的可能性，并返回[1,10]范围的随机数
			return num%10 + 1
		}
	}
}

// https://leetcode-cn.com/problems/implement-rand10-using-rand7/solution/cong-zui-ji-chu-de-jiang-qi-ru-he-zuo-dao-jun-yun-/

/*
已知 rand_N() 可以等概率的生成[1, N]范围的随机数
那么：
(rand_X() - 1) × Y + rand_Y() ==> 可以等概率的生成[1, X * Y]范围的随机数
即实现了 rand_XY()

*/
