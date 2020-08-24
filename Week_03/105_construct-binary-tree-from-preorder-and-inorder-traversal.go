package Week_03

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var inorderMap map[int]int
var pre []int

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap = make(map[int]int)
	for idx, num := range inorder {
		inorderMap[num] = idx
	}
	pre = preorder

	return helper(0, len(pre)-1, 0, len(pre)-1)
}
func helper(pl, pr, il, ir int) *TreeNode {

	if pl > pr {
		return nil
	}
	root := &TreeNode{
		Val: pre[pl],
	}
	idx := inorderMap[pre[pl]]
	root.Left = helper(pl+1, pl+idx-il, il, idx-1)
	root.Right = helper(pl+idx-il+1, pr, idx+1, ir)
	return root
}

func preTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	helper1(root, &ret)
	return ret
}
func helper1(root *TreeNode, ret *[]int) {

	if root == nil {
		return
	}
	*ret = append(*ret, root.Val)
	helper1(root.Left, ret)
	helper1(root.Right, ret)
}
