package Day74

import "strconv"

func decodeString(s string) string {
	sta := make(stack, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			tmp := make([]byte, 0)
			for sta.Peek() != '[' {
				tmp = append(tmp, sta.Pop())
			}
			sta.Pop() // pop '['
			cntStr := ""
			for !sta.IsEmpty() && sta.Peek() >= '0' && sta.Peek() <= '9' {
				cntStr = string(sta.Pop()) + cntStr
			}
			cnt, _ := strconv.Atoi(cntStr)
			for j := 0; j < cnt; j++ {
				for k := len(tmp) - 1; k >= 0; k-- {
					sta.Push(tmp[k])
				}
			}
		} else {
			sta.Push(s[i])
		}
	}
	return string(sta)
}

type stack []byte

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}
func (s *stack) Push(str byte) {
	*s = append(*s, str)
}
func (s *stack) Pop() byte {
	str := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return str
}
func (s *stack) Peek() byte {
	str := (*s)[s.Len()-1]
	return str
}
