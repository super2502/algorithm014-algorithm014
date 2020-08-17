package Week_02

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Right)
	right := invertTree(root.Left)

	root.Left = left
	root.Right = right
	return root
}
