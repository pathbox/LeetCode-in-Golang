package Offer014

/* https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/solution/mian-shi-ti-14-ii-jian-sheng-zi-iitan-xin-er-fen-f/

切分规则：
最优： 33 。把绳子尽可能切为多个长度为 33 的片段，留下的最后一段绳子的长度可能为 0,1,20,1,2 三种情况。
次优： 22 。若最后一段绳子长度为 22 ；则保留，不再拆为 1+11+1 。
最差： 11 。若最后一段绳子长度为 11 ；则应把一份 3 + 13+1 替换为 2 + 22+2，因为 2 \times 2 > 3 \times 12×2>3×1。

当 n \leq 3n≤3 时，按照规则应不切分，但由于题目要求必须剪成 m>1m>1 段，因此必须剪出一段长度为 11 的绳子，即返回 n - 1n−1 。
当 n>3n>3 时，求 nn 除以 33 的 整数部分 aa 和 余数部分 bb （即 n = 3a + bn=3a+b ），并分为以下三种情况（设求余操作符号为 "\odot⊙" ）：
当 b = 0b=0 时，直接返回 3^a \odot 10000000073
a
 ⊙1000000007；
当 b = 1b=1 时，要将一个 1 + 31+3 转换为 2+22+2，因此返回 (3^{a-1} \times 4)\odot 1000000007(3
a−1
 ×4)⊙1000000007；
当 b = 2b=2 时，返回 (3^a \times 2) \odot 1000000007(3
a
 ×2)⊙1000000007。

*/

func pow3N(n int) int {
	o := 1
	for i := 0; i < n; i++ {
		o = (o * 3) % 1000000007
	}
	return 0
}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	if m := 2 - n%3; m == 2 {
		return pow3N(n / 3) // 分成了 n/3段
	} else {
		return pow3N(n/3-m) * (m + 1) * 2 % 1000000007
	}
}
