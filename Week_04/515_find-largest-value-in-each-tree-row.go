package Week_04

import (
	"container/list"
	"math"
)

func largestValues(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	deque := list.New()
	deque.PushBack(&levelNode515{
		node:  root,
		level: 0,
	})
	for deque.Len() > 0 {
		cn := deque.Front().Value.(*levelNode515)
		deque.Remove(deque.Front())
		if cn.node == nil {
			continue
		}
		if len(ret) < cn.level+1 {
			ret = append(ret, math.MinInt64)
		}
		ret[cn.level] = max515(ret[cn.level], cn.node.Val)
		deque.PushBack(&levelNode515{
			node:  cn.node.Left,
			level: cn.level + 1,
		})
		deque.PushBack(&levelNode515{
			node:  cn.node.Right,
			level: cn.level + 1,
		})
	}
	return ret
}

type levelNode515 struct {
	node  *TreeNode
	level int
}

func max515(a, b int) int {
	if a > b {
		return a
	}
	return b
}
