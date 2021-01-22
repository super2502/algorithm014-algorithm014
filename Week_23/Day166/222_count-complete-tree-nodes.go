package Day166

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	height := 0
	start := root
	for start != nil {
		height++
		start = start.Left
	}
	left := pow2(height - 1)
	right := pow2(height) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if find(root, mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
func find(root *TreeNode, mid int) bool {
	str := fmt.Sprintf("%b", mid)
	parent := root
	for i := 0; i < len(str)-1; i++ {
		if str[i] == '0' {
			parent = parent.Left
		} else {
			parent = parent.Right
		}
	}
	if str[len(str)-1] == '0' {
		return parent.Left != nil
	} else {
		return parent.Right != nil
	}
}
func pow2(n int) int {
	return int(math.Pow(2.0, float64(n)))
}
