package Day127

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := make(stack, 0)
	s2 := make(stack, 0)
	for l1 != nil {
		s1.Push(l1)
		l1 = l1.Next
	}
	for l2 != nil {
		s2.Push(l2)
		l2 = l2.Next
	}
	var l3 *ListNode
	carry := 0
	for s1.Len() > 0 || s2.Len() > 0 {
		var a, b int
		if s1.Len() > 0 {
			a = s1.Pop().Val
		}
		if s2.Len() > 0 {
			b = s2.Pop().Val
		}
		sum := a + b + carry
		l3 = &ListNode{
			Val:  sum % 10,
			Next: l3,
		}
		carry = sum / 10
	}
	if carry > 0 {
		return &ListNode{
			Val:  carry,
			Next: l3,
		}
	}
	return l3
}

type stack []*ListNode

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Push(n *ListNode) {
	*s = append(*s, n)
}
func (s *stack) Pop() *ListNode {
	n := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return n
}
