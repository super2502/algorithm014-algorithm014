package Week_09

import "testing"

func Test_numDecodings(t *testing.T) {
	for str, way := range map[string]int{
		"12":   2,
		"226":  3,
		"0":    0,
		"1":    1,
		"2":    1,
		"1123": 5,
	} {
		ret := numDecodings(str)
		t.Logf("%v, expect %v, yours %v", str, way, ret)
	}
}
