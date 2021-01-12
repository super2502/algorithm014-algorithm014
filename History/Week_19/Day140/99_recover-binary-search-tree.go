package Day140

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var wrongLeft, wrongRight *TreeNode
	var dfs func(root *TreeNode)
	pre := math.MinInt64
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if root.Val < pre {
			wrongLeft = root
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)
	left := wrongLeft.Val
	var dfs1 func(root *TreeNode) bool
	dfs1 = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if dfs1(root.Left) {
			return true
		}
		if root.Val > left {
			wrongRight = root
			return true
		}
		return dfs1(root.Right)
	}
	dfs1(root)
	wrongLeft.Val, wrongRight.Val = wrongRight.Val, wrongLeft.Val
}
