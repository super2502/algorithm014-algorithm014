package Week_04

import "testing"

func Test_minMutation(t *testing.T) {

	ret := minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"})
	t.Logf("ret: %v", ret)
}
