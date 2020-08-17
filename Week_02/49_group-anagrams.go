package Week_02

import "sort"

func groupAnagrams(strs []string) [][]string {
	sortMap := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		key := string(s)
		//if _, ok := sortMap[key]; !ok {
		//	sortMap[key] = make([]string, 0)
		//}
		sortMap[key] = append(sortMap[key], str)
	}

	ret := make([][]string, 0, len(sortMap))
	for _, v := range sortMap {
		ret = append(ret, v)
	}
	return ret
}
