package Day29

import "container/list"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]struct{})
	visited := make(map[string]struct{})

	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}

	deque := list.New()
	deque.PushBack(beginWord)
	level := 1
	for deque.Len() > 0 {
		cns := make([]string, 0)
		for deque.Len() > 0 {
			cn := deque.Front().Value.(string)
			deque.Remove(deque.Front())
			cns = append(cns, cn)
		}
		for _, cn := range cns {
			nextWords := oneStep(cn, visited, wordMap)
			for word := range nextWords {
				if word == endWord {
					return level + 1
				} else {
					visited[word] = struct{}{}
					deque.PushBack(word)
				}
			}
		}
		level++
	}
	return 0
}

func oneStep(word string, visited map[string]struct{}, wordMap map[string]struct{}) map[string]struct{} {
	nextWords := make(map[string]struct{})
	for i := 0; i < len(word); i++ {
		for b := 'a'; b <= 'z'; b++ {
			key := word[:i] + string(b) + word[i+1:]
			if _, ok := visited[key]; ok {
				continue
			}
			if _, ok := wordMap[key]; ok {
				nextWords[key] = struct{}{}
			}
		}
	}

	return nextWords
}
