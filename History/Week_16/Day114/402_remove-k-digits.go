package Day114

func removeKdigits(num string, k int) string {
	s := make(stack, 0)
	for i := 0; i < len(num); i++ {
		for !s.isEmpty() && k > 0 && num[i] < s.peek() {
			s.pop()
			k--
		}
		s.push(num[i])
	}
	if k > 0 {
		s = s[:s.Len()-k]
	}
	for len(s) > 0 && s[0] == '0' {
		s = s[1:]
	}
	if len(s) == 0 {
		return "0"
	}
	return string(s)
}

type stack []byte

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) isEmpty() bool {
	return s.Len() == 0
}
func (s *stack) push(x byte) {
	*s = append(*s, x)
}
func (s *stack) pop() byte {
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x
}
func (s *stack) peek() byte {
	x := (*s)[s.Len()-1]
	return x
}
