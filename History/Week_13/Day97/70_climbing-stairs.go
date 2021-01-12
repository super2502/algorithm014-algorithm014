package Day97

func climbStairs(n int) int {
	first := 1
	if n == 1 {
		return first
	}
	second := 2
	for i := 3; i <= n; i++ {
		cur := first + second
		first = second
		second = cur
	}
	return second
}
