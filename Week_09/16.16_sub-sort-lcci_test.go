package Week_09

import (
	"testing"
)

func Test_subSort(t *testing.T) {
	nums := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	ret := []int{3, 9}

	r := subSort(nums)
	t.Logf("%+v, expect %v, yours %v", nums, ret, r)
}
