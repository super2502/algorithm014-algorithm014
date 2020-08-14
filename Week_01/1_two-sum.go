package Week_01

func twoSum(nums []int, target int) []int {

	l := len(nums)
	if l < 2 {
		return []int{}
	}
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		idx, ok := m[target-nums[i]]
		if ok {
			return []int{idx, i}
		}
		m[nums[i]] = i
	}

	return []int{}
}
