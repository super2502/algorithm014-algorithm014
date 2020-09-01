package Week_04

import "testing"

func Test_coinChange(t *testing.T) {
	ret := coinChange([]int{1, 2, 5}, 11)

	t.Logf("ret: %v", ret)
}
