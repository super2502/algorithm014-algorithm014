package Week_01

func trap(height []int) int {

	s := &stack{
		nodes: make([]*node, 0),
	}
	l := len(height)
	if l <= 2 {
		return 0
	}
	water := 0
	for i := 0; i < l; i++ {
		if s.isEmpty() {
			s.nodes = append(s.nodes, &node{
				idx: i,
				val: height[i],
			})
			continue
		}
		n := s.top()
		if n.val > height[i] {
			s.nodes = append(s.nodes, &node{
				idx: i,
				val: height[i],
			})
			continue
		}
		for !s.isEmpty() {
			n := s.pop()
			water += (i - n.idx) * (height[i] - n.val)
		}
	}

	return water
}

type stack struct {
	nodes []*node
}
type node struct {
	idx int
	val int
}

func (s *stack) isEmpty() bool {
	return len(s.nodes) == 0
}

func (s *stack) pop() *node {
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return node
}

func (s *stack) top() *node {
	if !s.isEmpty() {
		return s.nodes[len(s.nodes)-1]
	}
	return nil
}
