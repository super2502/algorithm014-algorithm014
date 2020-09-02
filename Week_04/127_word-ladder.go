package Week_04

import "container/list"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]struct{})
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	visited := make(map[string]bool)
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	deque := list.New()
	deque.PushBack(beginWord)
	l := 1
	for deque.Len() > 0 {
		cns := make([]string, 0)
		cnNext := make(map[string]struct{}, 0)
		for deque.Len() > 0 {
			cns = append(cns, deque.Front().Value.(string))
			deque.Remove(deque.Front())
		}
		for _, word := range cns {
			if word == endWord {
				return l + 1
			}
			nexts := nextWords(word, wordMap)
			for next := range nexts {
				if next == endWord {
					return l + 1
				}
				if _, ok := visited[next]; ok {
					continue
				}
				cnNext[next] = struct{}{}
			}
		}
		for next := range cnNext {
			deque.PushBack(next)
			visited[next] = true
		}
		l++
	}
	return 0
}

func nextWords(word string, wordMap map[string]struct{}) map[string]struct{} {
	ret := make(map[string]struct{})
	for i := 0; i < len(word); i++ {
		for b := 'a'; b <= 'z'; b++ {
			key := word[:i] + string(b) + word[i+1:]
			if _, ok := wordMap[key]; ok {
				ret[key] = struct{}{}
			}
		}
	}
	return ret
}
