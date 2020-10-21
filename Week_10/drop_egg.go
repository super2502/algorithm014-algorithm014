package Week_10

import (
	"fmt"
)

func dropEggs(level, cnt int) int {

	if level == 0 || cnt == 0 {
		return 0
	}
	eb := make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		eb[i] = make([]int, level)
	}
	for j := 0; j < level; j++ {
		eb[0][j] = j + 1
	}
	for i := 1; i < cnt; i++ {
		eb[i][0] = 1
		for j := 1; j < level; j++ {
			minCnt := level
			for k := 0; k <= j; k++ {
				//fmt.Printf("j: %v, k:%v, line(%+v)\n", j, k, eb[1])
				if k == 0 {
					minCnt = min(minCnt, 1+eb[i][j-k-1])
				} else if k == j {
					minCnt = min(minCnt, 1+eb[i-1][k-1])
				} else {
					minCnt = min(minCnt, 1+max(eb[i][j-k-1], eb[i-1][k-1]))
				}
			}
			eb[i][j] = minCnt
		}
	}
	for i := 0; i < cnt; i++ {
		fmt.Printf("1egg: %+v\n", eb[i])
	}
	return eb[cnt-1][level-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
