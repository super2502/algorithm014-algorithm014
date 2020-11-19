package Day101

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxRet := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxRet = max(maxRet, prices[i]-minPrice)
		}
	}
	return maxRet
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
