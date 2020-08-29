package Week_03

import (
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	x := []int{1, 1, 2}
	ret := permuteUnique(x)
	t.Logf("%+v", ret)
}
