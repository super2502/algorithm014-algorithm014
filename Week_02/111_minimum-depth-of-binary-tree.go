package Week_02

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 {
		return right + 1
	} else if right == 0 {
		return left + 1
	} else {
		return min(left, right) + 1
	}

}
