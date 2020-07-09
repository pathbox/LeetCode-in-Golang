package LeetCode121

/*
本题比较直接，只需要进行一次或者0次买卖操作即可。为啥会有0次？比如加个是[7,6,5,4,3,2,1]，你从哪天买都会亏本，还不如不买，最后受益0就是最大。

在这里我们需要两个变量，一个来保存截止目前为止，最低的价格是什么，还有一个变量保存截至目前为止，获得的最大受益是多少。比较直接，如果数组是[3,9,1,6,10],如果我们按照这种情况进行计算，9-3=6，当前最大是6，最低是3，接下来看1，如果我们要在价格是1的时候卖掉股票，必然在之前的最低点买入是最合适的，虽然这样亏钱，不过没关系，我们做了一个限制，如果价格比profit低，那就不更新，但是请注意，当前的1已经低于之前的最低点3了，那么更新最低点，接着往后找。

我们最终可以发现，1到10是最合适的价格，当然很多人看到这种股票，最先想到DP，本题也能用，但是有点浪费罢了
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
		if prices[i] < prices {
			minPrice = prices[i]
		}
	}
	return profit
}
