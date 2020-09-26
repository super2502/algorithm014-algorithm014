package Week_07

import (
	"strings"
)
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	l := len(beginWord)
	wordMap := make(map[string]bool)
	for _, word := range wordList {
		wordMap[word] = true
	}
	if !wordMap[beginWord] { wordMap[beginWord] = true }
	ret := make([][]string, 0)
	if !wordMap[endWord] {return ret}

	frontQueue := make(map[string][]string)
	backQueue := make(map[string][]string)
	frontQueue[beginWord] = []string{beginWord}
	backQueue[endWord] = []string{endWord}

	forward := true
	meet := false

	for len(frontQueue) > 0 {
		tmp := make(map[string][]string)
		for word, frontList := range frontQueue {
			for i := 0; i < l; i++ {
				for b:='a';b<='z';b++ {
					key := word[:i] + string(b) + word[i+1:]
					if backList, ok := backQueue[key]; ok {
						if forward {
							merge(&ret, frontList, backList)
						} else {
							merge(&ret, backList, frontList)
						}
						meet = true
					}
					if ! wordMap[key] {continue}

					for _, fwords := range frontList {
						if forward {
							tmp[key] = append(tmp[key], fwords+"-"+key)
						} else {
							tmp[key] = append(tmp[key], key + "-" + fwords)
						}
					}


				}
			}
		}
		if meet { break } // 找到最短路径了
		// 处理完本层所有节点再将其置为visited
		for key := range tmp {
			wordMap[key] = false
		}
		frontQueue = tmp
		if len(frontQueue) > len(backQueue) {
			frontQueue, backQueue = backQueue, frontQueue
			forward = ! forward
		}

	}

	return ret
}

func merge(ret *[][]string, frontList, backList []string) {
	for _, fwords := range frontList {
		for _, bwords := range backList {
			s := fwords + "-" + bwords
			sarr := strings.Split(s, "-")
			*ret = append(*ret, sarr)
		}
	}
}