package Week_03

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	first := 1
	second := 2
	var cur int
	for i := 3; i <= n; i++ {
		cur = first + second
		first = second
		second = cur
	}
	return cur
}
