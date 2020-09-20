package Week_07

type Trie struct {
	isEnd    bool
	children map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		children: make(map[byte]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, b := range []byte(word) {
		if _, ok := root.children[b]; !ok {
			root.children[b] = &Trie{
				children: make(map[byte]*Trie),
			}
		}
		root = root.children[b]
	}
	root.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for _, b := range []byte(word) {
		if _, ok := root.children[b]; !ok {
			return false
		}
		root = root.children[b]
	}
	return root.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for _, b := range []byte(prefix) {
		if _, ok := root.children[b]; !ok {
			return false
		}
		root = root.children[b]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
