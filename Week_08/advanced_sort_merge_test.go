package Week_08

import "testing"

func Test_mergeSort(t *testing.T) {
	for _, nums := range numss {
		mergeSort(nums, 0, len(nums)-1)
		t.Logf("%+v", nums)
	}
}
