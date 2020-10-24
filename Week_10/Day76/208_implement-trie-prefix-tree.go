package Day76

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		if root.children[idx] == nil {
			root.children[idx] = &Trie{}
		}
		root = root.children[idx]
	}
	root.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		if root.children[idx] == nil {
			return false
		}
		root = root.children[idx]
	}
	return root.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i := 0; i < len(prefix); i++ {
		idx := int(prefix[i] - 'a')
		if root.children[idx] == nil {
			return false
		}
		root = root.children[idx]
	}
	return true
}
