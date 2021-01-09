package Day151

func removeInvalidParentheses(s string) []string {
	queue := make(map[string]bool)
	ret := make([]string, 0)
	queue[s] = true
	for len(queue) > 0 {
		tmp := make(map[string]bool)
		for str := range queue {
			if validate(str) {
				ret = append(ret, str)
			}
			for i := 0; i < len(str); i++ {
				if str[i] != '(' && str[i] != ')' {
					continue
				}
				next := str[:i] + str[i+1:]
				//if i != len(str) - 1 {
				//	next += str[i+1:]
				//}
				tmp[next] = true
			}
		}
		if len(ret) > 0 {
			return ret
		}
		queue = tmp
	}
	return ret
}

func validate(str string) bool {
	s := make(stack, 0)
	for i := 0; i < len(str); i++ {
		if str[i] == ')' {
			if s.IsEmpty() {
				return false
			} else {
				s.Pop()
			}
		} else if str[i] == '(' {
			s.Push('(')
		}
	}
	if s.IsEmpty() {
		return true
	}
	return false
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
