package Day33

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	s1 := "abcde"
	s2 := "ace"
	ret := longestCommonSubsequence(s1, s2)
	t.Logf("%v", ret)
}
