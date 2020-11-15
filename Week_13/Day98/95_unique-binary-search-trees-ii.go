package Day98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var dfs func(i, j int) []*TreeNode
	dfs = func(i, j int) []*TreeNode {
		if i > j {
			return []*TreeNode{nil}
		}
		if i == j {
			tn := &TreeNode{
				Val: i,
			}
			return []*TreeNode{tn}
		}
		ret := make([]*TreeNode, 0)
		for k := i; k <= j; k++ {
			lefts := dfs(i, k-1)
			rights := dfs(k+1, j)
			for _, left := range lefts {
				for _, right := range rights {
					tn := &TreeNode{
						Val:   k,
						Left:  left,
						Right: right,
					}
					ret = append(ret, tn)
				}
			}
		}
		return ret
	}
	return dfs(1, n)
}
