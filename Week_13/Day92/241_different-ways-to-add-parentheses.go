package Day92

import (
	"strconv"
)

func diffWaysToCompute(input string) []int {
	nodes := make([]*node, 0)
	num := ""
	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			num += string(input[i])
		} else {
			val, _ := strconv.Atoi(num)
			nodes = append(nodes, &node{
				val: val,
				op:  input[i],
			})
			num = ""
		}
	}
	val, _ := strconv.Atoi(num)
	nodes = append(nodes, &node{
		val: val,
	})
	//fmt.Printf("nodes %+v\n", nodes)
	var dfs func(ns []*node) []int
	dfs = func(ns []*node) []int {
		if len(ns) == 0 {
			return []int{}
		}
		if len(ns) == 1 {
			return []int{ns[0].val}
		}
		ret := make([]int, 0)
		for i := 0; i < len(ns)-1; i++ {
			op := ns[i].op
			lefts := dfs(ns[:i+1])
			rights := dfs(ns[i+1:])
			for _, left := range lefts {
				for _, right := range rights {
					switch op {
					case '+':
						ret = append(ret, left+right)
					case '-':
						ret = append(ret, left-right)
					case '*':
						ret = append(ret, left*right)
					}
				}
			}
		}
		return ret
	}
	ans := dfs(nodes)
	return ans
}

type node struct {
	val int
	op  byte
}
