package Day66

import (
	"fmt"
	"strings"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ret := make([][]string, 0)
	pathWays := make([]string, 0)
	wordMap := make(map[string]bool)
	for _, word := range wordList {
		wordMap[word] = true
	}
	if !wordMap[endWord] {
		return ret
	}
	wordMap[endWord] = false
	frontMap := map[string][]string{beginWord: {beginWord}}
	backMap := map[string][]string{endWord: {endWord}}
	front := true
	for len(frontMap) != 0 {
		//fmt.Printf("======= round: %+v, %+v\n", frontMap, backMap)
		tmp := make(map[string][]string)
		for word, paths := range frontMap {
			nexts := nextWords(word, wordMap, backMap)
			fmt.Printf("nexts: %+v\n", nexts)
			for _, next := range nexts {
				otherPaths, ok := backMap[next]
				if ok {
					for _, path := range paths {
						for _, otherPath := range otherPaths {
							var s string
							if front {
								s = path + "-" + otherPath
							} else {
								s = otherPath + "-" + path
							}
							pathWays = append(pathWays, s)
						}
					}
				} else {
					if len(tmp[next]) == 0 {
						tmp[next] = make([]string, 0)
					}
					for _, path := range paths {
						if front {
							path = path + "-" + next
						} else {
							path = next + "-" + path
						}
						tmp[next] = append(tmp[next], path)
					}
				}
			}
		}
		if len(pathWays) > 0 {
			break
		}
		for visited := range tmp {
			wordMap[visited] = false
		}
		frontMap = tmp

		if len(frontMap) > len(backMap) {
			backMap, frontMap = frontMap, backMap
			front = !front
		}
	}
	for _, pathWay := range pathWays {
		arr := strings.Split(pathWay, "-")
		ret = append(ret, arr)
	}
	//fmt.Printf("wordMap: %+v\n", wordMap)
	//fmt.Printf("pathWays: %+v\n", pathWays)
	return ret
}

func nextWords(word string, wordMap map[string]bool, otherMap map[string][]string) []string {
	ret := make([]string, 0)
	for i := 0; i < len(word); i++ {
		for b := 'a'; b <= 'z'; b++ {
			key := word[:i] + string(b) + word[i+1:]
			if _, ok := otherMap[key]; ok {
				ret = append(ret, key)
			} else if wordMap[key] {
				ret = append(ret, key)
			}
		}
	}
	return ret
}
