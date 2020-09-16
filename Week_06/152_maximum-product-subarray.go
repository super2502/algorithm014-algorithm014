package Week_06

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	//dpMax := make([]int, n)
	//dpMin := make([]int, n)
	dpMax, dpMin := nums[0], nums[0]
	maxRet := nums[0]
	for i := 1; i < n; i++ {
		tmpMax, tmpMin := dpMax, dpMin
		if nums[i] < 0 {
			tmpMax, tmpMin = tmpMin, tmpMax
		}
		dpMax = max152(tmpMax*nums[i], nums[i])
		dpMin = min152(tmpMin*nums[i], nums[i])
		maxRet = max152(dpMax, maxRet)
	}

	return maxRet
}

func maxProduct1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0], dpMin[0] = nums[0], nums[0]
	maxRet := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			dpMax[i] = max152(dpMin[i-1]*nums[i], nums[i])
			dpMin[i] = min152(dpMax[i-1]*nums[i], nums[i])
		} else {
			dpMax[i] = max152(dpMax[i-1]*nums[i], nums[i])
			dpMin[i] = min152(dpMin[i-1]*nums[i], nums[i])
		}
		maxRet = max152(dpMax[i], maxRet)
	}

	return maxRet
}

func max152(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min152(a, b int) int {
	if a < b {
		return a
	}
	return b
}
