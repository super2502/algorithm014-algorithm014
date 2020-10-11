package Day63

func longestValidParentheses(s string) int {
	maxLength := 0
	st := make(stack, 0)
	st.Push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st.Push(i)
		} else {
			idx := st.Pop()
			if idx == -1 || s[idx] != '(' {
				st.Push(i)
			} else {
				maxLength = max(maxLength, i-st.Peek())
			}
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}
func (s *stack) Push(x int) {
	*s = append(*s, x)
}
func (s *stack) Pop() int {
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x
}
func (s *stack) Peek() int {
	x := (*s)[s.Len()-1]
	return x
}
