package Day83

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret int

func distributeCoins(root *TreeNode) int {
	ret = 0
	helper(root)
	return ret
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	right := helper(root.Right)
	ret += abs(left) + abs(right)
	return root.Val + left + right - 1
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
