package Week_02

func postorderTraversal(root *TreeNode) []int {
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
			s.push(n)
			s.push(&ColorNode{
				Node:  n.Node.Right,
				Color: white,
			})
			s.push(&ColorNode{
				Node:  n.Node.Left,
				Color: white,
			})
		} else {
			ret = append(ret, n.Node.Val)
		}
	}
	return ret
}
