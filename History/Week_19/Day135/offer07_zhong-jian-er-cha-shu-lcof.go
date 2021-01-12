package Day135

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for idx, val := range inorder {
		inorderMap[val] = idx
	}
	var helper func(l, r, l1, r1 int) *TreeNode
	helper = func(l, r, l1, r1 int) *TreeNode {
		if l > r {
			return nil
		}
		node := &TreeNode{
			Val: preorder[l],
		}
		inIdx := inorderMap[preorder[l]]
		node.Left = helper(l+1, inIdx-l1+l, l1, inIdx-1)
		node.Right = helper(inIdx-l1+l+1, r, inIdx+1, r1)
		return node
	}
	return helper(0, len(preorder)-1, 0, len(inorder)-1)
}
