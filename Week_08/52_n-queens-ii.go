package Week_08

import "math/bits"

var cnt int

func totalNQueens(n int) int {
	cnt = 0
	dfs52([]int{}, n, 0, 0, 0)
	return cnt
}

func dfs52(queens []int, n int, col, ld, rd int) {
	if len(queens) == n {
		cnt++
		return
	}
	bitMap := ^(col | ld | rd) & ((1 << n) - 1)
	for bitMap > 0 {
		i := bitMap & (-bitMap)
		bitMap = bitMap & (bitMap - 1)
		q := bits.OnesCount(uint(i - 1))
		queens = append(queens, q)
		dfs52(queens, n, col|i, (ld|i)<<1, (rd|i)>>1)
		queens = queens[:len(queens)-1]
	}
}
