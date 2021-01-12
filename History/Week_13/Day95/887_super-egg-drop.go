package Day95

import (
	"fmt"
	"math"
)

func superEggDrop(K int, N int) int {

	dp := make([][]int, K)
	dp[0] = make([]int, N+1)
	for j := 0; j <= N; j++ {
		dp[0][j] = j
	}
	for i := 1; i < K; i++ {
		dp[i] = make([]int, N+1)
		dp[i][1] = 1
		for j := 2; j <= N; j++ {
			minCnt := math.MaxInt64
			l, r := 1, j
			for l <= r {
				mid := l + (r-l)>>1
				broken := dp[i-1][mid-1]
				nonBroken := dp[i][j-mid]
				//fmt.Printf("l, r, mid (%v)(%v)(%v), b, nb (%v)(%v)\n", l, r , mid,broken, nonBroken)
				if broken == nonBroken {
					minCnt = min(minCnt, broken)
					break
				}
				if broken > nonBroken {
					minCnt = min(minCnt, broken)
					r = mid - 1
				} else {
					minCnt = min(minCnt, nonBroken)
					l = mid + 1
				}
			}
			//fmt.Printf("==== i,j (%v)(%v)got min (%v)\n",i ,j ,minCnt)
			dp[i][j] = minCnt + 1
		}
	}
	for i := 0; i < K; i++ {
		fmt.Printf("%+v\n", dp[i])
	}
	return dp[K-1][N]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
