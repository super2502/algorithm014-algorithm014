package Day21

func getKthMagicNumber(k int) int {
	ret := make([]int, k)
	ret[0] = 1
	i, j, l := 0, 0, 0

	for p := 1; p < k; p++ {
		m := min(ret[i]*3, min(ret[j]*5, ret[l]*7))
		if m == ret[i]*3 {
			i++
		}
		if m == ret[j]*5 {
			j++
		}
		if m == ret[l]*7 {
			l++
		}
		ret[p] = m

	}
	return ret[k-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
