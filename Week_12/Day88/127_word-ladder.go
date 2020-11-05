package Day88

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool)
	for _, word := range wordList {
		wordMap[word] = true
	}
	if !wordMap[endWord] {
		return 0
	}
	delete(wordMap, endWord)
	frontMap := map[string]bool{beginWord: true}
	backMap := map[string]bool{endWord: true}
	l := 1
	for len(frontMap) > 0 {
		l++
		tmp := make(map[string]bool)
		for word := range frontMap {
			nexts := toNext(word, wordMap, backMap)
			for next := range nexts {
				if backMap[next] {
					return l
				}
				tmp[next] = true
			}
		}
		frontMap = tmp
		if len(frontMap) > len(backMap) {
			frontMap, backMap = backMap, frontMap
		}
	}
	return 0
}

func toNext(word string, wordMap, backMap map[string]bool) map[string]bool {
	ret := make(map[string]bool)
	for i := 0; i < len(word); i++ {
		for b := 'a'; b <= 'z'; b++ {
			key := word[:i] + string(b) + word[i+1:]
			if wordMap[key] || backMap[key] {
				ret[key] = true
				delete(wordMap, key)
			}
		}
	}
	return ret
}
