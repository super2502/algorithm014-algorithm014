package Day79_107

func findTheDifference(s string, t string) byte {
	sm := [26]int{}
	for i := 0; i < len(s); i++ {
		sm[s[i]-'a']++
		sm[t[i]-'a']--
	}
	sm[t[len(s)]-'a']--
	for b, cnt := range sm {
		if cnt == -1 {
			return byte(b + 'a')
		}
	}
	return ' '
}
