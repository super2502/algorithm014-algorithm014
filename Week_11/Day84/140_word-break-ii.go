package Day84

import (
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	tr := &trie{}
	for _, word := range wordDict {
		tr.insert(word)
	}
	memo := make(map[int][][]string)

	var dfs func(idx int) [][]string
	dfs = func(idx int) [][]string {
		if memo[idx] != nil {
			return memo[idx]
		}
		ret := make([][]string, 0)
		if idx == len(s) {
			return ret
		}
		for i := idx; i < len(s); i++ {
			word := s[idx : i+1]
			if !tr.beginWith(word) {
				break
			}
			if tr.search(word) {
				//fmt.Printf("try word(%v)\n", word)
				for _, nexts := range dfs(i + 1) {
					//fmt.Printf("get %+v\n", nexts)
					ret = append(ret, append([]string{word}, nexts...))
				}
				if i+1 == len(s) {
					ret = append(ret, []string{word})
				}
			}
		}
		//fmt.Printf("ret with idx(%v) (%+v)\n", idx, ret)
		memo[idx] = ret
		return ret
	}
	ret := dfs(0)
	ans := make([]string, 0)
	for _, path := range ret {
		ans = append(ans, strings.Join(path, " "))
	}
	return ans
}

type trie struct {
	isWord   bool
	children [26]*trie
}

func (tr *trie) insert(word string) {
	root := tr
	for i := 0; i < len(word); i++ {
		key := int(word[i] - 'a')
		if root.children[key] == nil {
			root.children[key] = &trie{}
		}
		root = root.children[key]
	}
	root.isWord = true
}
func (tr *trie) search(word string) bool {
	root := tr
	for i := 0; i < len(word); i++ {
		key := int(word[i] - 'a')
		if root.children[key] == nil {
			return false
		}
		root = root.children[key]
	}
	return root.isWord
}

func (tr *trie) beginWith(word string) bool {
	root := tr
	for i := 0; i < len(word); i++ {
		key := int(word[i] - 'a')
		if root.children[key] == nil {
			return false
		}
		root = root.children[key]
	}
	return true
}
