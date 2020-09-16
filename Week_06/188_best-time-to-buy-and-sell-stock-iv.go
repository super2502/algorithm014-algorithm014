package Week_06

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	ldp := make([][]int, 2)
	ldp[0] = make([]int, n)
	ldp[1] = make([]int, n)
	ldp[0][0] = 0
	ldp[1][0] = -1 * prices[0]

	maxP := 0
	for j := 1; j <= k; j++ {
		cdp := make([][]int, 2)
		cdp[0] = make([]int, n)
		cdp[1] = make([]int, n)
		cdp[0][0] = 0
		cdp[1][0] = -1 * prices[0]

		for i := 1; i < n; i++ {
			cdp[0][i] = max188(cdp[0][i-1], ldp[1][i-1]+prices[i])
			cdp[1][i] = max188(cdp[1][i-1], cdp[0][i-1]-prices[i])
		}
		maxP = max188(maxP, cdp[0][n-1])
		//ldp = cdp
		deepCopy(cdp, ldp)
	}
	return maxP
}

func max188(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func deepCopy(src, dest [][]int) {
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src[0]); j++ {
			dest[i][j] = src[i][j]
		}
	}
}
