package Day77

type WordDictionary struct {
	isWord   bool
	children [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		key := int(word[i] - 'a')
		if root.children[key] == nil {
			root.children[key] = &WordDictionary{}
		}
		root = root.children[key]
	}
	root.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.isWord
	}
	if word[0] != '.' {
		key := int(word[0] - 'a')
		if this.children[key] == nil {
			return false
		}
		return this.children[key].Search(word[1:])
	}
	for i := 0; i < 26; i++ {
		if this.children[i] != nil {
			found := this.children[i].Search(word[1:])
			if found {
				return true
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
