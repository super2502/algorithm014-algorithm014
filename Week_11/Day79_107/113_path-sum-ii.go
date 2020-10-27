package Day79_107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	ret := make([][]int, 0)
	dfs([]int{}, &ret, root, sum)
	return ret
}

func dfs(path []int, ret *[][]int, root *TreeNode, sum int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			p := make([]int, len(path))
			copy(p, path)
			*ret = append(*ret, p)
		}
		return
	}
	dfs(path, ret, root.Left, sum-root.Val)
	dfs(path, ret, root.Right, sum-root.Val)
	path = path[:len(path)-1]
}
