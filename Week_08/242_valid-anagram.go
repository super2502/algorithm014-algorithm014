package Week_08

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	rs := []rune(s)
	rt := []rune(t)
	m := make(map[rune]int)
	for i := 0; i < len(rs); i++ {
		m[rs[i]]++
		m[rt[i]]--
	}
	for _, cnt := range m {
		if cnt != 0 {
			return false
		}
	}
	return true
}
