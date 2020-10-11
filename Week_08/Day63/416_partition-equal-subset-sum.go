package Day63

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	m := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum&1 == 1 {
		return false
	}
	half := sum >> 1

	return dfs(nums, 0, half, m)
}

func dfs(nums []int, start int, target int, m map[int]map[int]bool) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	if start == len(nums) {
		return false
	}
	if _, ok := m[start]; !ok {
		m[start] = make(map[int]bool)
	}
	var ok1, ok2 bool
	if !m[start][target-nums[start]] {
		m[start][target-nums[start]] = true
		ok1 = dfs(nums, start+1, target-nums[start], m)
	}
	if !m[start][target] {
		m[start][target] = true
		ok2 = dfs(nums, start+1, target, m)
	}

	return ok1 || ok2
}
