package Day63

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	for _, s := range []string{
		"",
		"(",
		")",
		"(((",
		"()()",
		"(()()",
		"((())",
		")(()))",
		")()(()",
		"())()()",
	} {
		maxLength := longestValidParentheses(s)
		t.Logf("%s, %v", s, maxLength)
	}

}
