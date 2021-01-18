package Day160

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	math.MinInt64
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		ret = append(ret, queue[0].Val)
		tmp := make([]*TreeNode, 0)
		for _, node := range queue {
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
		}
		queue = tmp
	}
	return ret
}
