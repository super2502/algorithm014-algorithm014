package Day67

import (
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	nums := []int{3, 1, 2, 1}
	ret := permuteUnique(nums)
	for _, row := range ret {
		t.Logf("%+v", row)
	}
}
