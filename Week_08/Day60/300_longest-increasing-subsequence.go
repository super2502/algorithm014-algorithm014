package Day60

func lengthOfLIS(nums []int) int {
	dp := make([]int, 0, len(nums))
	binarySearch := func(res *[]int, target int) {
		l, r := 0, len(*res)-1
		for l <= r {
			mid := l + (r-l)>>1
			if (*res)[mid] == target {
				return
			} else if (*res)[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if r == len(*res)-1 {
			*res = append(*res, target)
		} else {
			if (*res)[r+1] > target {
				(*res)[r+1] = target
			}
		}
	}
	for _, num := range nums {
		binarySearch(&dp, num)
		// fmt.Printf("dp: %+v\n", dp)
	}
	return len(dp)
}
func lengthOfLIS1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		maxLenI := 0
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] && maxLenI < dp[j] {
				maxLenI = dp[j]
			}
		}
		dp[i] = 1 + maxLenI

	}
	maxLen := 0
	for i := 0; i < n; i++ {
		if maxLen < dp[i] {
			maxLen = dp[i]
		}
	}
	return maxLen
}
