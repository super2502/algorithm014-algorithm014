package Week_06

import (
	"fmt"
	"math"
)

//[
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}
	dp := make([]int, m)
	dp[0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min120(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	ret := math.MaxInt64
	for i := 0; i < m; i++ {
		ret = min120(ret, dp[i])
	}
	return ret
}

func minimumTotal1(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}
	dp := make([][]int, m)
	dp[0] = make([]int, 1)
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		//fmt.Printf("== %+v %v, %v, %v\n", i, dp[i][0], dp[i-1][0], triangle[i][0])
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
		for j := 1; j < i; j++ {
			dp[i][j] = min120(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			fmt.Printf("== %v, %v, %v, %v, %v\n", i, j, dp[i-1][j-1], dp[i-1][j], triangle[i][j])
		}
		fmt.Printf("%+v\n", dp[i])
	}
	ret := math.MaxInt64
	for i := 0; i < m; i++ {
		ret = min120(ret, dp[m-1][i])
	}
	return ret
}

func min120(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
