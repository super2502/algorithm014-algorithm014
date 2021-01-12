package Day141

func reverseParentheses(s string) string {
	st := make(stack, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			tmp := make([]byte, 0)
			for !st.IsEmpty() {
				b := st.Pop()
				if b != '(' {
					tmp = append(tmp, b)
				} else {
					break
				}
			}
			for j := 0; j < len(tmp); j++ {
				st.Push(tmp[j])
			}
		} else {
			st.Push(s[i])
		}
	}
	return string(st)
}

type stack []byte

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}
func (s *stack) Push(b byte) {
	*s = append(*s, b)
}
func (s *stack) Pop() byte {
	b := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return b
}
func (s *stack) Peek() byte {
	b := (*s)[s.Len()-1]

	return b
}
