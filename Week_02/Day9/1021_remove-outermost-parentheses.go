package Day9

func removeOuterParentheses(S string) string {
	l := len(S)
	c := 0
	ret := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		if c == 0 {
			c++
			continue
		} // 栈起点左括号抛弃
		if S[i] == '(' {
			c++ // 非起点左括号入栈，追加到结果
			ret = append(ret, S[i])
			continue
		}
		if S[i] == ')' {
			c--
		} // 遇到右括号出栈一个左括号
		if c != 0 {
			ret = append(ret, S[i])
		} // 非栈终点右括号追加到结果
	}
	return string(ret)
}
