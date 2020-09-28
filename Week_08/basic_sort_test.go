package Week_08

import (
	"testing"
)

var numss = [][]int {
	{1,3,5,6,8,3,2,4},
	{},
	{9,4,6,3,8,1,2},
}

func TestInsertionSort(t *testing.T) {
	for _, nums := range numss {
		t.Logf("%+v", insertionSort(nums))
	}
}

func TestSelectionSort(t *testing.T) {
	for _, nums := range numss {
		t.Logf("%+v", selectionSort(nums))
	}
}

func TestBubbleSort(t *testing.T) {
	for _, nums := range numss {
		t.Logf("%+v", bubbleSort(nums))
	}
}