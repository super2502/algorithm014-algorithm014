package Week_09

import "testing"

func Test_racecar(t *testing.T) {
	for _, ca := range [][]int{
		{3, 2},
		{6, 5},
		{5, 7},
	} {
		ret := racecar(ca[0])
		t.Logf("target: %v, expect: %v, yours: %v", ca[0], ca[1], ret)
	}
}
