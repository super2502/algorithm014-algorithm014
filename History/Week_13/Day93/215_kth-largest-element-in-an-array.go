package Day93

func findKthLargest(nums []int, k int) int {
	ret := quickSelect(nums, k, 0, len(nums)-1)
	//fmt.Printf("qs nums (%+v)\n", nums)
	return ret
}

func quickSelect(nums []int, k int, l, r int) int {
	//pivot := rand.Int() % ( r - l + 1) + l
	//nums[pivot], nums[r] = nums[r], nums[pivot]
	idx := l
	for i := l; i < r; i++ {
		if nums[i] > nums[r] {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}
	nums[idx], nums[r] = nums[r], nums[idx]
	if idx == k-1 {
		return nums[idx]
	}
	if idx < k-1 {
		return quickSelect(nums, k, idx+1, r)
	} else {
		return quickSelect(nums, k, l, idx-1)
	}
}
