package Week_04

import "container/list"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ret := make([][]string, 0)
	wordMap := make(map[string]struct{})
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	visited := make(map[string]struct{})
	if _, ok := wordMap[endWord]; !ok {
		return ret
	}
	deque := list.New()
	deque.PushBack(&levelNode126{
		word: beginWord,
		path: []string{beginWord},
	})
	for deque.Len() > 0 {
		cns := make([]*levelNode126, 0)
		for deque.Len() > 0 {
			cns = append(cns, deque.Front().Value.(*levelNode126))
			deque.Remove(deque.Front())
		}
		nextIgnore := make(map[string]struct{})
		for _, cn := range cns {
			nextWords := oneStep(cn.word, visited, wordMap)
			for next := range nextWords {
				path := make([]string, len(cn.path))
				copy(path, cn.path)
				path = append(path, next)
				if next == endWord {
					ret = append(ret, path)
				} else {
					deque.PushBack(&levelNode126{
						word: next,
						path: path,
					})
				}
				nextIgnore[next] = struct{}{}
			}
		}
		for next := range nextIgnore {
			visited[next] = struct{}{}
		}
		if len(ret) > 0 {
			return ret
		}
	}

	return ret
}

type levelNode126 struct {
	word string
	path []string
}

func oneStep(word string, visited map[string]struct{}, dictMap map[string]struct{}) map[string]struct{} {
	ret := make(map[string]struct{})
	for i := 0; i < len(word); i++ {
		for b := 'a'; b <= 'z'; b++ {
			key := word[:i] + string(b) + word[i+1:]
			if _, ok := visited[key]; ok {
				continue
			}
			_, ok := dictMap[key]
			if ok {
				ret[key] = struct{}{}
			}
		}
	}
	return ret
}
