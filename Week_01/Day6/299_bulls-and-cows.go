package Day6

import "fmt"

func getHint(secret string, guess string) string {
	m := make([]int, 10)
	bulls := 0
	cows := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
			continue
		}
		if (m[secret[i]-'0'] + 1) <= 0 {
			cows++
		}
		if (m[guess[i]-'0'] - 1) >= 0 {
			cows++
		}
		m[secret[i]-'0']++
		m[guess[i]-'0']--
	}
	return fmt.Sprintf("%vA%vB", bulls, cows)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
