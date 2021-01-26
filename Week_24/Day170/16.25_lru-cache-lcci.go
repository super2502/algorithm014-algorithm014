package Day170

type node struct {
	key  int
	val  int
	pre  *node
	next *node
}
type LRUCache struct {
	capacity int
	length   int
	root     *node
	cache    map[int]*node
}

func Constructor(capacity int) LRUCache {
	r := LRUCache{
		capacity: capacity,
		length:   0,
		cache:    make(map[int]*node),
	}
	root := &node{}
	root.next = root
	root.pre = root
	r.root = root
	return r
}

func (this *LRUCache) isFull() bool {
	return this.capacity == this.length
}
func (this *LRUCache) isEmpty() bool {
	return this.length == 0
}
func (this *LRUCache) front() *node {
	if this.isEmpty() {
		return nil
	}
	return this.root.next
}
func (this *LRUCache) back() *node {
	if this.isEmpty() {
		return nil
	}
	return this.root.pre
}
func (this *LRUCache) pop(n *node) {
	pre, next := n.pre, n.next
	pre.next = next
	next.pre = pre
	n.pre, n.next = nil, nil
}
func (this *LRUCache) pushFront(n *node) {
	next := this.root.next
	this.root.next, next.pre = n, n
	n.pre, n.next = this.root, next
}
func (this *LRUCache) moveToFront(n *node) {
	this.pop(n)
	this.pushFront(n)
}
func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToFront(n)
	return n.val
}
func (this *LRUCache) Put(key int, value int) {
	n, ok := this.cache[key]
	if ok {
		n.key = key
		n.val = value
		this.moveToFront(n)
	} else {
		if this.isFull() {
			end := this.back()
			this.pop(end)
			delete(this.cache, end.key)
			this.length--
		}
		n := &node{key: key, val: value}
		this.pushFront(n)
		this.cache[key] = n
		this.length++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
