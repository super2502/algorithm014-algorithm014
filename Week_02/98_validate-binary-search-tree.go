package Week_02

import "math"

var pre int

func isValidBST(root *TreeNode) bool {
	pre = math.MinInt64

	return helper98(root)
}

func helper98(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !helper98(root.Left) {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return helper98(root.Right)
}
