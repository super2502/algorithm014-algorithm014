package Day157

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	mark = 1000
)

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	markDepth(root)
	var ret *TreeNode
	for root != nil {
		if root.Left == nil && root.Right == nil {
			ret = root
			break
		} else if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			a, b := root.Left.Val/mark, root.Right.Val/mark
			if a == b {
				ret = root
				break
			} else if a > b {
				root = root.Left
			} else {
				root = root.Right
			}
		}
	}
	restore(ret)
	return ret
}

func markDepth(root *TreeNode) {
	if root == nil {
		return
	}
	markDepth(root.Left)
	markDepth(root.Right)
	var a, b int
	if root.Left != nil {
		a = root.Left.Val / mark
	}
	if root.Right != nil {
		b = root.Right.Val / mark
	}
	root.Val = root.Val + mark*(max(a, b)+1)
}
func restore(root *TreeNode) {
	if root == nil {
		return
	}
	root.Val = root.Val % mark
	restore(root.Left)
	restore(root.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
