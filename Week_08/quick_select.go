package Week_08

func quickSelectKthMin(nums []int, k int) (int, int, int) {
	// 0 < k <= len(nums)
	n := len(nums)
	start, end := 0, n
	swaps := 0
	loops := 0
	for {
		loops++
		pivot := end - 1
		idx := start
		for i := start; i < end - 1; i++ {
			if nums[i] < nums[pivot] {
				nums[i], nums[idx] = nums[idx], nums[i]
				swaps++
				idx++
			}
		}
		nums[idx], nums[pivot] = nums[pivot], nums[idx]
		if idx == k - 1 {
			return loops, swaps, nums[idx]
		} else if idx < k - 1 {
			start = idx + 1
		} else {
			end = idx
		}
	}
}

func quickSelectKthMax(nums []int, k int) (int, int, int) {
	// 0 < k <= len(nums)
	n := len(nums)
	start, end := 0, n
	swaps := 0
	loops := 0
	for {
		loops++
		pivot := end - 1
		idx := start
		for i := start; i < end - 1; i++ {
			if nums[i] > nums[pivot] {
				nums[i], nums[idx] = nums[idx], nums[i]
				swaps++
				idx++
			}
		}
		nums[idx], nums[pivot] = nums[pivot], nums[idx]
		if idx == k - 1 {
			return loops, swaps, nums[idx]
		} else if idx < k - 1 {
			start = idx + 1
		} else {
			end = idx
		}
	}
}