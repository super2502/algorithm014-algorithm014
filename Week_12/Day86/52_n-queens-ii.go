package Day86

func totalNQueens(n int) int {
	var ret int
	var dfs func(row, n int, col, ld, rd int)
	dfs = func(row, n int, col, ld, rd int) {
		if row == n {
			ret++
			return
		}
		bitMap := ^(col | ld | rd) & ((1 << n) - 1)
		for bitMap > 0 {
			pos := bitMap & (-bitMap)
			bitMap = bitMap & (bitMap - 1)
			dfs(row+1, n, col|pos, (ld|pos)<<1, (rd|pos)>>1)
		}
	}
	dfs(0, n, 0, 0, 0)
	return ret
}
