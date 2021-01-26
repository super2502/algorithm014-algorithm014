package Day168

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	ret := 0
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		if node.Left != nil && node.Left.Val == node.Val {
		} else {
			left = 0
		}
		if node.Right != nil && node.Right.Val == node.Val {
		} else {
			right = 0
		}
		ret = max(ret, left+right)
		return 1 + max(left, right)
	}
	helper(root)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
