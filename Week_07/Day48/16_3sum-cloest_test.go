package Day48

import "testing"

func Test3SumClosest(t *testing.T) {
	nums := []int{-1,2,1,-4}
	target := 1
	t.Logf("%+v, %v", nums, target)
	ret := threeSumClosest(nums, target)
	t.Logf("ret: %v", ret)
}
