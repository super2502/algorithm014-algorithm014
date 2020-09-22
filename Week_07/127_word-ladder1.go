package Week_07

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]struct{})
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	length := 1
	front := map[string]struct{}{beginWord: {}}
	back := map[string]struct{}{endWord: {}}
	for len(front) > 0 {
		length += 1
		tmp := make(map[string]struct{})
		for word := range front {
			for i := 0; i < len(word); i++ {
				for b := 'a'; b <= 'z'; b++ {
					key := word[:i] + string(b) + word[i+1:]
					if _, ok := back[key]; ok {
						return length
					}
					if _, ok := wordMap[key]; ok {
						tmp[key] = struct{}{}
						delete(wordMap, key)
					}
				}
			}
		}
		front = tmp
		if len(front) > len(back) {
			front, back = back, front
		}
	}
	return 0
}
