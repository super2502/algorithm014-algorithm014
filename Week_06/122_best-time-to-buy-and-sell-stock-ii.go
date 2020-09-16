package Week_06

func maxProfit122(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	sum := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}
