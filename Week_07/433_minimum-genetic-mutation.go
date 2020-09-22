package Week_07

func minMutation(start string, end string, bank []string) int {
	bankMap := make(map[string]struct{})
	for _, gin := range bank {
		bankMap[gin] = struct{}{}
	}
	if _, ok := bankMap[end]; !ok {
		return -1
	}
	front := map[string]struct{}{start: {}}
	back := map[string]struct{}{end: {}}
	count := 0
	for len(front) > 0 {
		count++
		tmp := make(map[string]struct{})
		for word := range front {
			for i := 0; i < 8; i++ {
				for _, b := range []byte{'A', 'C', 'G', 'T'} {
					key := word[:i] + string(b) + word[i+1:]
					// fmt.Printf("key %s (%v)(%T)\n", key, string(b), b)
					if _, ok := back[key]; ok {
						return count
					}
					if _, ok := bankMap[key]; ok {
						tmp[key] = struct{}{}
						delete(bankMap, key)
					}
				}
			}
		}
		front = tmp
		if len(front) > len(back) {
			front, back = back, front
		}
	}
	return -1
}
