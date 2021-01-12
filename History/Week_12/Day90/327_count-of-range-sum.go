package Day90

func countRangeSum(nums []int, lower int, upper int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum >= lower && sum <= upper {
			cnt++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum >= lower && sum <= upper {
				cnt++
			}
		}
	}
	return cnt
}
