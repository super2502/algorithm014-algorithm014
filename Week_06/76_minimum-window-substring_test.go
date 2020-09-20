package Week_06

import "testing"

func Test_minWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	//s = "AEBANC"
	s = "AAA"
	t1 := "ABC"
	x := minWindow(s, t1)
	t.Logf("%v", x)
}
