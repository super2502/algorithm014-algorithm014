package Week_02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ColorNode struct {
	Node  *TreeNode
	Color int
}

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	white, gary := 0, 1
	s := make(stack, 0)
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
			s.push(n)
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

type stack []*ColorNode

func (s *stack) len() int {
	return len(*s)
}
func (s *stack) isEmpty() bool {
	return s.len() == 0
}
func (s *stack) push(n *ColorNode) {
	*s = append(*s, n)
}

func (s *stack) pop() *ColorNode {
	n := (*s)[s.len()-1]
	*s = (*s)[:s.len()-1]
	return n
}

func inorderTraversal1(root *TreeNode) []int {
	ret := make([]int, 0)

	helper(root, &ret)

	return ret
}

func helper(root *TreeNode, ret *[]int) {

	if root == nil {
		return
	}
	helper(root.Left, ret)
	*ret = append(*ret, root.Val)
	helper(root.Right, ret)
}
