package Day100

import (
	"math"
)

func maxProduct(nums []int) int {
	maxP, minP := 1, 1
	maxRet := math.MinInt64
	for _, num := range nums {
		if num > 0 {
			maxP, minP = max(num, maxP*num), min(num, minP*num)
		} else {
			maxP, minP = max(num, minP*num), min(num, maxP*num)
		}
		maxRet = max(maxRet, maxP)
		//fmt.Printf("num(%v), maxP(%v), minP(%v), maxRet(%v)\n", num, maxP, minP, maxRet)
	}
	return maxRet
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
