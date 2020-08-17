package Week_02

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	l := len(s)
	m := make(map[byte]int)

	for i := 0; i < l; i++ {
		m[s[i]]++
		m[t[i]]--
	}

	for _, count := range m {
		if count != 0 {
			return false
		}
	}

	return true
}
