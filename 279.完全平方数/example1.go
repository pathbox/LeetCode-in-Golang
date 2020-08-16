package LeetCode279

// 使用暴力枚举法会超出时间限制的原因很简单，因为我们重复的计算了中间解
/*
DP动态规划：
1.用切片存小于n的所有平方数
2.切片dp存0-n的n+1个数的完全平方和的最小个数
3.从1开始向n依次求各个数的完全平方数和的最小个数。
4.dp[1]是最基本的等式，后面的所有要求数的完全平方数的个数都是围绕dp[1]=1计算。dp[0]置零
5.求第i个数的最小平方数个数时，以7，8，9为例：
dp[7]=min(dp[7-1]+1,dp[7-4]+1) ==>min(4,4)==>dp[7]=4
dp[8]=min(dp[8-1]+1,dp[8-4]+1) ==>min(4,2)==>dp[8]=2
dp[9]=min(dp[9-1]+1,dp[9-4]+1,dp[9-9]+1) ==>min(2,3,1)==>dp[9]=1
*/
func numSquares(n int) int {
	//把平方数写入切片
	sum := []int{}
	for i := 0; i <= n/2+2; i++ {
		sum = append(sum, i*i)
	}
	//dp存放0-n的最小平方数个数
	dp := make([]int, 0, 1000)
	dp = append(dp, 0)
	//第一个循环求0-n的最小平方数个数
	for i := 1; i <= n; i++ {
		mid, p, j := n, n, 1
		// 第二个循环求i的最小平方数个数
		for ; i >= sum[j]; j++ {
			p = min(dp[i-sum[j]]+1, mid)
			mid = p
		}
		dp = append(dp, mid)
	}
	return dp[n]
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
