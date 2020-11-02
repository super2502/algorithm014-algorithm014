package Day85

import "math/bits"

func solveNQueens(n int) [][]string {
	ret := make([][]string, 0)
	var dfs func(path []int, n int, col, ld, rd int)
	dfs = func(path []int, n int, col, ld, rd int) {
		if len(path) == n {
			printQ(path, &ret)
			return
		}
		bitMap := ^(col | ld | rd) & ((1 << n) - 1)
		for bitMap > 0 {
			p := bitMap & (-bitMap)
			bitMap = bitMap & (bitMap - 1)
			i := bits.OnesCount(uint(p - 1))
			path = append(path, i)
			dfs(path, n, col|p, (ld|p)<<1, (rd|p)>>1)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, n, 0, 0, 0)
	return ret
}

func printQ(path []int, ret *[][]string) {
	ss := make([]string, len(path))
	for i := 0; i < len(path); i++ {
		bs := make([]byte, len(path))
		for j := 0; j < len(path); j++ {
			bs[j] = '.'
		}
		bs[path[i]] = 'Q'
		ss[i] = string(bs)
	}
	*ret = append(*ret, ss)
}
