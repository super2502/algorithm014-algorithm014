package Week_09

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	ret := 4

	r := lengthOfLIS(nums)
	t.Logf("%+v, expect %v, yours %v", nums, ret, r)
}
