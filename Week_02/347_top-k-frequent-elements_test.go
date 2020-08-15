package Week_02

import (
	"testing"
)

func TestTopKFrequent(t *testing.T) {

	var nums = []int{1, 1, 1, 2, 2, 3}
	k := 2

	ret := topKFrequent(nums, k)

	t.Logf("ret (%+v)", ret)
}
