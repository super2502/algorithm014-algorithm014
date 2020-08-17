package Week_02

type ColorNode2 struct {
	Node  *Node
	Color int
}

func preorder(root *Node) []int {
	ret := make([]int, 0)
	white, gary := 0, 1
	s := make(stack2, 0)
	s.push(&ColorNode2{
		Node: root,
	})

	for !s.isEmpty() {
		n := s.pop()
		if n.Node == nil {
			continue
		}
		if n.Color == white {
			for i := len(n.Node.Children) - 1; i >= 0; i-- {
				s.push(&ColorNode2{
					Node: n.Node.Children[i],
				})

			}
			n.Color = gary
			s.push(n)
		} else {
			ret = append(ret, n.Node.Val)
		}
	}

	return ret

}

type stack2 []*ColorNode2

func (s *stack2) len() int {
	return len(*s)
}
func (s *stack2) isEmpty() bool {
	return s.len() == 0
}
func (s *stack2) push(n *ColorNode2) {
	*s = append(*s, n)
}
func (s *stack2) pop() *ColorNode2 {
	n := (*s)[s.len()-1]
	*s = (*s)[:s.len()-1]
	return n
}
