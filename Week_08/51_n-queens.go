package Week_08

import (
	"math/bits"
)

var ret [][]string

func solveNQueens(n int) [][]string {
	ret = make([][]string, 0)
	dfs51([]int{}, n, 0, 0, 0)
	return ret
}

func dfs51(queens []int, n int, col, ld, rd int) {
	if len(queens) == n {
		printQueens(queens)
		return
	}
	//fmt.Printf("%08b\n", col)
	//fmt.Printf("%08b\n", ld)
	//fmt.Printf("%08b\n", rd)
	bitMap := ^(col | ld | rd) & ((1 << n) - 1) // 高位清零
	//fmt.Printf("%08b ======\n", bitMap)
	for bitMap > 0 {
		i := bitMap & (-bitMap)          // 得到最低位1， 即为本层皇后可放置的一种可能位置
		bitMap = bitMap & (bitMap - 1)   // 干掉最低位1
		q := bits.OnesCount(uint(i - 1)) // 2的q次幂 减1后,1的个数就是q...
		queens = append(queens, q)
		dfs51(queens, n, col|i, (ld|i)<<1, (rd|i)>>1) // drill down
		queens = queens[:len(queens)-1]               // revert
	}
}

func printQueens(queens []int) {
	//fmt.Printf("queens: %+v\n", queens)
	n := len(queens)
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		bs := make([]byte, n)
		for j := 0; j < n; j++ {
			bs[j] = '.'
		}
		bs[queens[i]] = 'Q'
		ss[i] = string(bs)
	}
	ret = append(ret, ss)
}
