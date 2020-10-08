package Week_08

import "testing"

func Test_isUnique(t *testing.T) {
	s := "laab"
	//s = "abc"
	s = "leetcode"
	s = "leeb"
	t.Logf("ret: %v", isUnique(s))
}
