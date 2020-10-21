package Day73

func isValid(s string) bool {
	m := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if b, ok := m[s[i]]; !ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != b {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
