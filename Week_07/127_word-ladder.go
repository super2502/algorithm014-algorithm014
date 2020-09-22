package Week_07

import (
	"container/list"
)

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
	length := 1
	for deque.Len() > 0 {
		num := deque.Len()
		for i := 0; i < num; i++ {
			cn := deque.Front().Value.(string)
			deque.Remove(deque.Front())
			nexts := nextWords(cn, wordMap, visited)
			for next := range nexts {
				if next == endWord {
					return length + 1
				}
				deque.PushBack(next)
			}
		}
		length++
	}
	return 0
}

func nextWords(word string, wordMap map[string]struct{}, visited map[string]struct{}) map[string]struct{} {
	ret := make(map[string]struct{})
	for i := 0; i < len(word); i++ {
		for j := 'a'; j <= 'z'; j++ {
			key := word[:i] + string(j) + word[i+1:]
			if _, ok := wordMap[key]; !ok {
				continue
			}
			if _, ok := visited[key]; ok {
				continue
			}
			ret[key] = struct{}{}
			visited[key] = struct{}{}
		}
	}
	return ret
}
