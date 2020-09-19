package Week_06

func countSubstrings(s string) int {

	count := 0
	for i := 0; i<len(s);i++ {
		count += is(s, i, i)
		if i != len(s) - 1 {
			count += is(s, i, i+1)
		}
	}
	return count
}

func is(s string, i, j int) int {
	count := 0
	for i >= 0 && j < len(s) {
		if i == j  {
			count++
		} else if s[i] == s[j] {
			count++
		} else {
			break
		}
		i--
		j++
	}
	return count
}


