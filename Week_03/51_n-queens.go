package Week_03

import "fmt"

// 模板
// path: 每行的Queen的位置数组， len(path)==n即可返回
// 列表： 1-n行
// 不在列表的条件： 竖，斜

var ret [][]int

func solveNQueens(n int) [][]string {
	ret = make([][]int, 0)
	path := make([]int, 0)

	dfs(path, n)

	fmt.Printf(" ====== %v 皇后一共有 %v种方法 ====\n", n, len(ret))
	answer := make([][]string, len(ret))
	for i := 0; i < len(ret); i++ {
		answer[i] = make([]string, 0)
		for j := 0; j < n; j++ {
			s := ""
			for k := 0; k < n; k++ {
				if k == ret[i][j] {
					s += "Q"
				} else {
					s += "."
				}
			}
			answer[i] = append(answer[i], s)
		}
	}
	return answer
}

func dfs(path []int, n int) {
	if len(path) == n {
		//fmt.Printf("path: %+v\n", path)
		p := make([]int, len(path))
		copy(p, path)
		ret = append(ret, p)
	}

	for i := 0; i < n; i++ {
		var cannot bool
		for line, q := range path {
			if q == i {
				cannot = true
				break
			}
			if q-i == len(path)-line || i-q == len(path)-line {
				cannot = true
				break
			}
		}
		if cannot {
			continue
		}

		path = append(path, i)
		dfs(path, n)
		path = path[:len(path)-1]
	}
}
