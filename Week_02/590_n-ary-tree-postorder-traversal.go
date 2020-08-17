package Week_02

///**
// * Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

type ColorNode1 struct {
	Node  *Node
	Color int
}

type stack1 []*ColorNode1

func (s *stack1) len() int {
	return len(*s)
}
func (s *stack1) isEmpty() bool {
	return s.len() == 0
}
func (s *stack1) push(n *ColorNode1) {
	*s = append(*s, n)
}

func (s *stack1) pop() *ColorNode1 {
	n := (*s)[s.len()-1]
	*s = (*s)[:s.len()-1]
	return n
}

func postorder(root *Node) []int {
	ret := make([]int, 0)
	s := make(stack1, 0)
	white, gary := 0, 1
	s.push(&ColorNode1{
		Node: root,
	})
	for !s.isEmpty() {
		n := s.pop()
		if n.Node == nil {
			continue
		}

		if n.Color == white {
			n.Color = gary
			s.push(n)
			for i := len(n.Node.Children) - 1; i >= 0; i-- {
				s.push(&ColorNode1{
					Node: n.Node.Children[i],
				})
			}
		} else {
			ret = append(ret, n.Node.Val)
		}
	}
	return ret
}

func postorder1(root *Node) []int {
	ret := make([]int, 0)
	helper590(root, &ret)
	return ret
}

func helper590(root *Node, ret *[]int) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		helper590(child, ret)
	}
	*ret = append(*ret, root.Val)
}
