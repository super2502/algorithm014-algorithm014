package Week_08

import "container/list"

type LRUCache struct {
	deque    *list.List
	capacity int
	cache    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		deque:    list.New(),
		capacity: capacity,
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	ele, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.deque.MoveToFront(ele)
	return ele.Value.(*node).v
}

func (this *LRUCache) Put(key int, value int) {
	ele, ok := this.cache[key]
	if ok {
		ele.Value.(*node).v = value
		this.deque.MoveToFront(ele)
		return
	}
	if this.deque.Len() == this.capacity {
		ele := this.deque.Back()
		this.deque.Remove(ele)
		delete(this.cache, ele.Value.(*node).k)
	}
	this.deque.PushFront(&node{k: key, v: value})
	this.cache[key] = this.deque.Front()
}

type node struct {
	k int
	v int
}
