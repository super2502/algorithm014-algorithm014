package Day30

func subarraySum(nums []int, k int) int {

	m := make(map[int]int)
	m[0] = 1
	count := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if c, ok := m[sum-k]; ok {
			count += c
		}
		m[sum]++
	}
	return count
}
