package Week_11

func smallerNumbersThanCurrent(nums []int) []int {
	ret := make([]int, len(nums))
	if len(nums) == 0 {
		return ret
	}
	idxArr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		idxArr[i] = i
	}
	quickSort(nums, idxArr, 0, len(nums)-1)
	ret[idxArr[0]] = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			ret[idxArr[i]] = ret[idxArr[i-1]]
		} else {
			ret[idxArr[i]] = i
		}
	}
	return ret
}
func quickSort(nums, idxArr []int, i, j int) {
	if i >= j {
		return
	}
	idx, pivot := i, j
	for k := i; k < j; k++ {
		if nums[k] < nums[pivot] {
			nums[idx], nums[k] = nums[k], nums[idx]
			idxArr[idx], idxArr[k] = idxArr[k], idxArr[idx]
			idx++
		}
	}
	nums[pivot], nums[idx] = nums[idx], nums[pivot]
	idxArr[pivot], idxArr[idx] = idxArr[idx], idxArr[pivot]
	quickSort(nums, idxArr, i, idx-1)
	quickSort(nums, idxArr, idx+1, j)
}
