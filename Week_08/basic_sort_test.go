package Week_08

import (
	"testing"
)

var numss = [][]int{
	{1, 3, 5, 6, 8, 3, 2, 4},
	{},
	{9, 4, 6, 3, 8, 1, 2},
	{1, 2, 3, 4, 5, 6, 7, 8},
}

func TestInsertionSort(t *testing.T) {
	for _, nums := range numss {
		insertionSort(nums)
		t.Logf("%+v", nums)
	}
}

func TestSelectionSort(t *testing.T) {
	for _, nums := range numss {
		selectionSort(nums)
		t.Logf("%+v", nums)
	}
}

func TestBubbleSort(t *testing.T) {
	for _, nums := range numss {
		bubbleSort(nums)
		t.Logf("%+v", nums)
	}
}
