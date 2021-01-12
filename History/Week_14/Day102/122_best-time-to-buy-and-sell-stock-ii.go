package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	ret := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ret += prices[i] - prices[i-1]
		}
	}
	return ret
}
