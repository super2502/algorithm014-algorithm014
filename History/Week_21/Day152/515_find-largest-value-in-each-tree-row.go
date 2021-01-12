package Day152

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmp := make([]*TreeNode, 0)
		max := math.MinInt64
		for _, node := range queue {
			if max < node.Val {
				max = node.Val
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		ret = append(ret, max)
		queue = tmp
	}
	return ret
}
