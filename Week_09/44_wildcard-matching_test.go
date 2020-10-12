package Week_09

import "testing"

func Test_isMatch(t *testing.T) {
	cases := [][]string{
		{"aa", "a"},
		{"aa", "*"},
		{"cb", "?a"},
		{"adceb", "*a*b"},
		{"acdcb", "a*c?b"},
	}
	for _, row := range cases {
		ok := isMatch(row[0], row[1])
		t.Logf("%v match %v, %v", row[0], row[1], ok)
	}
}
