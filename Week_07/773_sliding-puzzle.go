package Week_07

import "strconv"

func slidingPuzzle(board [][]int) int {
	distMap := map[int][]int{
		0: {1, 3},
		1: {0, 2, 4},
		2: {1, 5},
		3: {0, 4},
		4: {1, 3, 5},
		5: {2, 4},
	}
	var start string
	startIdx := 0
	for i := 0; i < 6; i++ {
		num := board[i/3][i%3]
		if num == 0 {
			startIdx = i
		}
		start += strconv.Itoa(num)
	}
	end := "123450"
	visited := make(map[string]bool)
	queue := make(map[string]*node)
	queue[start] = &node{
		s:   start,
		idx: startIdx,
	}
	count := 0
	for len(queue) > 0 {
		tmp := make(map[string]*node)
		for s, cn := range queue {
			if s == end {
				return count
			}
			for _, dist := range distMap[cn.idx] {
				next := toNext(s, cn.idx, dist)
				if visited[next] {
					continue
				}
				tmp[next] = &node{
					s:   next,
					idx: dist,
				}
				visited[next] = true
			}
		}
		queue = tmp
		count++
	}
	return -1
}
func toNext(s string, idx, dist int) string {
	bs := []byte(s)
	bs[idx], bs[dist] = bs[dist], bs[idx]
	return string(bs)
}

type node struct {
	s   string
	idx int
}
