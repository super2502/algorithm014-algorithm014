package Day96

func minCostClimbingStairs(cost []int) int {
	first, second := 0, 0
	cur := min(first, cost[0])
	for i := 1; i < len(cost); i++ {
		cur = min(first+cost[i-1], second+cost[i])
		first = second
		second = cur
	}
	return cur
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
