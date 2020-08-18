package Week_02

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	inOrderIdx := findIdx(inorder, rootVal)

	leftPre := preorder[1 : inOrderIdx+1]
	leftIn := inorder[0:inOrderIdx]
	rightPre := preorder[inOrderIdx+1 : len(preorder)]
	rightIn := inorder[inOrderIdx+1 : len(preorder)]

	n := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(leftPre, leftIn),
		Right: buildTree(rightPre, rightIn),
	}
	return n
}

func findIdx(nums []int, x int) int {
	for idx, val := range nums {
		if val == x {
			return idx
		}
	}
	return 0
}
