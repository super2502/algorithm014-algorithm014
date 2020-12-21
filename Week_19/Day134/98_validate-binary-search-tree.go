package Day134

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	ok, _ := helper(root, pre)
	return ok
}

func helper(root *TreeNode, pre int) (bool, int) {
	if root == nil {
		return true, pre
	}
	var ok bool
	ok, pre = helper(root.Left, pre)
	if !ok {
		return false, pre
	}
	if root.Val <= pre {
		return false, pre
	}
	pre = root.Val
	return helper(root.Right, pre)
}
