package Week_02

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	s := make(stack, 0)
	white, gary := 0, 1
	s.push(&ColorNode{
		Node:  root,
		Color: white,
	})
	for !s.isEmpty() {
		n := s.pop()
		if n.Node == nil {
			continue
		}
		if n.Color == white {
			n.Color = gary
			s.push(&ColorNode{
				Node:  n.Node.Right,
				Color: white,
			})
			s.push(&ColorNode{
				Node:  n.Node.Left,
				Color: white,
			})
			s.push(n)

		} else {
			ret = append(ret, n.Node.Val)
		}
	}
	return ret
}

func preorderTraversal1(root *TreeNode) []int {
	ret := make([]int, 0)

	helper1(root, &ret)

	return ret
}
func helper1(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.Val)
	helper(root.Left, ret)
	helper(root.Right, ret)
}
