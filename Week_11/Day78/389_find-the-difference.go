package Day78

func findTheDifference(s string, t string) byte {
	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	m[t[len(s)]-'a']--
	for idx, cnt := range m {
		if cnt == -1 {
			return byte(idx + 'a')
		}
	}
	return ' '
}
