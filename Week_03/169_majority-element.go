package Week_03

func majorityElement(nums []int) int {
	m := make(map[int]int)

	maxNum := 0
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > maxCount {
			maxNum = nums[i]
			maxCount = m[nums[i]]
		}
	}
	return maxNum
}
