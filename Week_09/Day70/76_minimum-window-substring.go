package Day70

func minWindow(s string, t string) string {
	ret := ""
	if len(s) == 0 || len(t) == 0 {
		return ret
	}
	tMap := [128]int{}
	counter := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
		counter[t[i]]++
	}
	initMap := [128]int{}
	j := 0
	minLen := len(s) + 1
	cnt := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if tMap[b] == 0 {
			continue
		}
		initMap[b]++
		if initMap[b] == tMap[b] {
			cnt++
		}
		for cnt == len(counter) {
			if tMap[s[j]] == 0 {
				j++
				continue
			}
			if minLen > i-j+1 {
				minLen = i - j + 1
				ret = s[j : i+1]
			}
			initMap[s[j]]--
			if initMap[s[j]] < tMap[s[j]] {
				cnt--
			}
			j++
		}
	}

	return ret
}
