package Day44

type Trie struct {
	IsEnd    bool
	Children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, b := range []byte(word) {
		if root.Children[b-'a'] == nil {
			root.Children[b-'a'] = &Trie{}
		}
		root = root.Children[b-'a']
	}
	root.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for _, b := range []byte(word) {
		if root.Children[b-'a'] == nil {
			return false
		}
		root = root.Children[b-'a']
	}
	return root.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for _, b := range []byte(prefix) {
		if root.Children[b-'a'] == nil {
			return false
		}
		root = root.Children[b-'a']
	}
	return true
}
