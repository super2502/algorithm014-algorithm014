package Week_06

import (
	"math"
)
func splitArray(nums []int, m int) int {
	left, right := max(nums), sum(nums)
	// fmt.Printf("left(%v), right (%v)\n",left, right)
	for left < right {
		mid := left + ( right - left) / 2
		cnt := countNoBiggerThenK(nums ,mid)
		// fmt.Printf("cnt/m (%v)(%v), mid (%v)\n",cnt, m, mid)
		if cnt <= m {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func countNoBiggerThenK(nums []int, k int) int {
	count := 1
	sum := 0
	for i := 0; i< len(nums); i++ {
		sum += nums[i]
		if sum > k {
			count++
			sum = nums[i]
		}
	}
	return count
}

func sum(nums []int) int {

	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}

func max(nums []int) int {
	m := math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}