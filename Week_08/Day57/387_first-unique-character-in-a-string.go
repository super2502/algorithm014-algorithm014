package Day57

func firstUniqChar(s string) int {
	charMap := make(map[rune]int, len(s))
	for _, b := range s {
		charMap[b]++
	}
	for idx, b := range s {
		if charMap[b] == 1 {
			return idx
		}
	}
	return -1
}
