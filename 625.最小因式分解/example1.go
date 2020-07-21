package LeetCode625

/*
解题思路：
求的是分解之后组合最小，例如18，可以分级为2，9； 3；6
1、小于10的数，不用分解，就是本身最小
2、分集常规从9开始遍历到2尝试，比如18，求得是9，2 组合就是92，记得逆序即可
3、另外记得18因式分解到1，就是18 * 1 不符合题意得，所以结束条件是 >1
*/
func smallestFactorization(a int) int {
	if a < 10 {
		return a
	}
	var res int
	for i := 9; i >= 2; i-- { // 例如18求的得结果是92，需要反转
		for a%i == 0 && a > 1 {
			res = res*10 + i
			a = a / i
		}
	}
	if a > 10 {
		return 0
	}
	return helper(res)
}

// helper 实现整数反转得能力
func helper(num int) int {
	var res int
	for num > 0 {
		res = res*10 + num%10
		num = num / 10
	}
	if res > 1<<31 { // 根据题中要求加了过滤
		return 0
	}
	return res
}
