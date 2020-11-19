package main

import (
	"testing"
)

func TestWa(t *testing.T) {

	nums := []int{4,2,3,3,5,2}

	ret := maxWa(nums)
	t.Logf("%+v, ret: %v", nums, ret)


}

func TestMaxArea(t *testing.T) {

	nums := []int{4,2,3,3,5,2}
	nums = []int{5,4,3,2,1}

	ret := maxArea(nums)
	t.Logf("%+v, ret: %v", nums, ret)


}

func TestPerfect(t *testing.T) {
	recs := [][]int{
		{1,1,3,3},
		{3,1,4,2},
		{3,2,4,4},
		{1,3,2,4},
		{2,3,3,4},
	}
	recs = [][]int{
		{1,1,3,3},
		{3,1,4,2},
		{3,2,4,4},
		{1,3,2,4},
		//{2,3,3,4},
	}
	recs = [][]int{
		{1,1,3,3},
		{3,1,4,2},
		//{3,2,4,4},
		{1,3,2,4},
		{2,2,4,4},
	}
	ret := perfect(recs)
	t.Logf("ret: %v", ret)

}