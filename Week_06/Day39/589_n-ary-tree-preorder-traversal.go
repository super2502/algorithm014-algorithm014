package Day39

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	ret := make([]int, 0)
	s := make(stack, 0)
	if root == nil {
		return ret
	}
	s.Push(root)
	for !s.isEmpty() {
		cn := s.Pop()
		ret = append(ret, cn.Val)
		for i := len(cn.Children) - 1; i >= 0; i-- {
			s.Push(cn.Children[i])
		}
	}
	return ret
}

type stack []*Node

func (s *stack) len() int {
	return len(*s)
}
func (s *stack) isEmpty() bool {
	return s.len() == 0
}
func (s *stack) Push(x *Node) {
	*s = append(*s, x)
}
func (s *stack) Pop() *Node {
	x := (*s)[s.len()-1]
	*s = (*s)[:s.len()-1]
	return x
}
func (s *stack) Peek() *Node {
	x := (*s)[s.len()-1]
	return x
}
