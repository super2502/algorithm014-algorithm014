package Week_02

func nthUglyNumber(n int) int {

	i, j, k := 0, 0, 0
	ret := make([]int, n)
	ret[0] = 1
	for p := 1; p < n; p++ {
		m := min264(ret[i]*2, min264(ret[j]*3, ret[k]*5))
		if m == ret[i]*2 {
			i++
		}
		if m == ret[j]*3 {
			j++
		}
		if m == ret[k]*5 {
			k++
		}

		ret[p] = m
	}
	return ret[n-1]
}

func min264(a, b int) int {
	if a < b {
		return a
	}
	return b
}
