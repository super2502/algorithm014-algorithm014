package Day72

func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n < 2 {
		return []int{}
	}
	cache := make(map[int]int)
	for i := 0; i < n; i++ {
		if idx, ok := cache[target-nums[i]]; ok {
			return []int{idx, i}
		}
		cache[nums[i]] = i
	}
	return []int{}
}
