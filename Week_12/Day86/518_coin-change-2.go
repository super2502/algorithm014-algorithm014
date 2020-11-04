package Day86

import (
	"fmt"
	"math"
)

func change(amount int, coins []int) int {

	store := make(map[int]int)
	for idx := range coins {
		store[idx] = int(math.Pow(10, float64(idx+1)))
	}

	if amount == 0 {
		return 1
	}
	dp := make([]map[int]struct{}, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make(map[int]struct{})
		for idx, coin := range coins {
			fmt.Printf("debug ??? (%v)(%v)(%v)\n", idx, coin, i)
			if i-coin < 0 {
				continue
			}
			if i-coin == 0 {
				dp[i][store[idx]] = struct{}{}
			} else {
				for key := range dp[i-coin] {
					dp[i][key+store[idx]] = struct{}{}
				}
			}
			fmt.Printf("dp map len(%v) (+%v) i (%v)\n", len(dp[i]), dp[i], i)
		}
	}
	return len(dp[amount])
}
