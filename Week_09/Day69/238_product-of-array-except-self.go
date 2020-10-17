package Day69

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = 1
	}
	left, right := 1, 1
	i, j := 0, n-1
	for i < n {
		ret[i] *= left
		ret[j] *= right
		left *= nums[i]
		right *= nums[j]
		i++
		j--
	}
	return ret
}
func productExceptSelf1(nums []int) []int {
	n := len(nums)
	prefixProduct := make([]int, n)
	suffixProduct := make([]int, n)
	ret := make([]int, n)
	prefixProduct[0] = nums[0]
	suffixProduct[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		prefixProduct[i] = nums[i] * prefixProduct[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		suffixProduct[i] = nums[i] * suffixProduct[i+1]
	}
	ret[0] = suffixProduct[1]
	ret[n-1] = prefixProduct[n-2]
	for i := 1; i < n-1; i++ {
		ret[i] = prefixProduct[i-1] * suffixProduct[i+1]
	}

	return ret
}
