package Week_09

func countSmaller(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	if n == 0 {
		return ret
	}
	ret[n-1] = 0
	for i := n - 1; i >= 0; i-- {
		cnt := 0
		j := i + 1
		for j < n {
			if nums[i] > nums[j] {
				cnt = ret[j] + 1
				break
			} else if nums[i] == nums[j] {
				cnt = ret[j]
				break
			} else {
				j++
			}
		}
		ret[i] = cnt
	}
	return ret
}
