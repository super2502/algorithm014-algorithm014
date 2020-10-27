package Day79

import "strings"

func wordPattern(pattern string, s string) bool {
	pm := make(map[byte]string)
	sm := make(map[string]byte)
	sarr := strings.Split(s, " ")
	if len(pattern) != len(sarr) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if str, ok := pm[pattern[i]]; ok {
			if str != sarr[i] {
				return false
			}
		} else {
			pm[pattern[i]] = sarr[i]
		}
		if b, ok := sm[sarr[i]]; ok {
			if b != pattern[i] {
				return false
			}
		} else {
			sm[sarr[i]] = pattern[i]
		}
	}
	return true
}
