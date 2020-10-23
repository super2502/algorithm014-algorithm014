package Day75

import "container/list"

type LRUCache struct {
	capacity int
	list     *list.List
	hash     map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		hash:     make(map[int]*list.Element),
	}
}
func (this *LRUCache) Get(key int) int {
	ele, ok := this.hash[key]
	if !ok {
		return -1
	}
	this.list.MoveToFront(ele)
	return ele.Value.(*node).val
}
func (this *LRUCache) Put(key int, value int) {
	ele, ok := this.hash[key]
	if ok {
		ele.Value.(*node).val = value
		this.list.MoveToFront(ele)
	} else {
		this.list.PushFront(&node{
			key: key,
			val: value,
		})
		this.hash[key] = this.list.Front()
		if this.isFull() {
			rmEle := this.list.Back()
			this.list.Remove(rmEle)
			delete(this.hash, rmEle.Value.(*node).key)
		}
	}
}
func (this *LRUCache) isFull() bool {
	return this.list.Len() > this.capacity
}

type node struct {
	key int
	val int
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
