package Week_09

import "testing"

func Test_isMatch10(t *testing.T) {
	cases := [][]string{
		{"aaa", "ab*ac*a"},
		{"mississippi", "mis*is*p*."},
		{"aa", "*a"},
	}
	for _, row := range cases {
		ok := isMatch10(row[0], row[1])
		t.Logf("%v match %v, %v", row[0], row[1], ok)
	}
}
