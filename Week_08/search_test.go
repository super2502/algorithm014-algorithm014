package Week_08

import (
	"testing"
)

func TestFirstLessEq(t *testing.T) {

	nums := []int{1,4,6,6,7}
	t.Logf("%v", nums)
	t.Logf("============= testng last less eq")
	for _, target := range []int{
		0,1,2,3,4,5,6,7,8,
	} {
		var i int
		i = searchLastLessEq(nums, target)
		t.Logf("target: %v, idx: %v", target, i)
	}
	t.Logf("============= testng first large eq")
	for _, target := range []int{
		0,1,2,3,4,5,6,7,8,
	} {
		var i int
		i = searchFirstLargeEq(nums, target)
		t.Logf("target: %v, idx: %v", target, i)
	}

}

func TestFirstLastEq(t *testing.T) {
	nums := []int{1,4,6,6,6,7}
	nums = []int{0,1,2,3,4,5,6,6,6,7}
	t.Logf("%v", nums)
	t.Logf("============= testng first eq")
	for _, target := range []int{
		6,
	} {
		var i int
		i = searchFirstEq(nums, target)
		t.Logf("target: %v, idx: %v", target, i)
	}

	t.Logf("============= testng last eq")
	for _, target := range []int{
		6,
	} {
		var i int
		i = searchLastEq(nums, target)
		t.Logf("target: %v, idx: %v", target, i)
	}

}

