package Week_02

import "sort"

func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)
	l := len(strs)
	if l == 0 {
		return ret
	}
	m := make(map[string][]string)
	for i := 0; i < l; i++ {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		key := string(s)
		if _, ok := m[key]; !ok {
			m[key] = make([]string, 0)
		}
		m[key] = append(m[key], strs[i])
	}

	for _, words := range m {
		ret = append(ret, words)
	}
	return ret
}
