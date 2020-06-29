package Offer063

func maxProfit(prices []int) int {
	minPrice := 0
	record := make([]int, len(prices)) // 第i天的最高利润

	if len(prices) == 0 {
		return 0
	} else {
		minPrice = prices[0]
		record[0] = 0
	}

	for i := 1; i < len(record); i++ {
		recordBefore := record[i-1]
		recordNow := prices[i] - minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if recordNow > recordBefore {
			record[i] = recordNow
		} else {
			record[i] = recordBefore
		}
	}
	return record[len(record)-1]
}
