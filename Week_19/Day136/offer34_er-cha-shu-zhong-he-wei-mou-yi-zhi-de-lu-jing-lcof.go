package Day136

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	var dfs func(path []int, sum int, root *TreeNode)
	dfs = func(path []int, sum int, root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			if root.Val == sum {
				p := make([]int, len(path))
				copy(p, path)
				p = append(p, root.Val)
				ret = append(ret, p)
			}
			return
		}
		path = append(path, root.Val)
		if root.Left != nil {
			dfs(path, sum-root.Val, root.Left)
		}
		if root.Right != nil {
			dfs(path, sum-root.Val, root.Right)
		}
		path = path[:len(path)-1]
	}
	dfs([]int{}, sum, root)
	return ret
}
