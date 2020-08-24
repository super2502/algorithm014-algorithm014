package Day15

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	helper(root, &ret)
	return ret
}
func helper(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, ret)
	*ret = append(*ret, root.Val)
	helper(root.Right, ret)
}
