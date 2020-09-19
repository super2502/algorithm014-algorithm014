package Week_06

func longestValidParentheses(s string) int {

	n := len(s)
	if n == 0 {
		return 0
	}

	maxLen := 0
	st := make(stack, 0)
	st.Push(-1)
	for i:= 0; i< n;i++ {
		if s[i] == '(' {
			st.Push(i)
		} else {
			st.Pop()
			if !st.IsEmpty() {
				maxLen = max32(maxLen, i-st.Peek())
			} else {
				st.Push(i)
			}
		}
	}
	return maxLen
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
	x := (*s)[s.Len() - 1]
	*s = (*s)[:s.Len() - 1]
	return x
}
func (s *stack) Peek() int {
	x := (*s)[s.Len() - 1]
	return x
}

func max32(a, b int) int {
	if a > b {
		return a
	}
	return b
}