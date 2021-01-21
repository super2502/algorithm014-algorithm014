package Day165

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		ret = max(ret, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
