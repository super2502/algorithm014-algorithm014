package Week_01

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	l        *list.List
	m        map[int]*list.Element
	capacity int
}

type kv struct {
	key int
	val int
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{
		l:        list.New(),
		m:        make(map[int]*list.Element),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	ele, ok := this.m[key]
	if !ok {
		return -1
	}
	this.l.MoveToFront(ele)
	return ele.Value.(*kv).val
}

func (this *LRUCache) Put(key int, value int) {
	ele, ok := this.m[key]
	if ok {
		ele.Value.(*kv).val = value
		this.l.MoveToFront(ele)
		return
	}
	if this.isFull() {
		//fmt.Printf("list len (%v)\n", this.l.Len())
		rmEle := this.l.Back()

		this.l.Remove(rmEle)
		delete(this.m, rmEle.Value.(*kv).key)
	}

	this.l.PushFront(&kv{
		key: key,
		val: value,
	})
	this.m[key] = this.l.Front()

	//this.print()
}
func (this *LRUCache) print() {
	for k, v := range this.m {
		fmt.Printf("cache: %v, %v, %T, \n", k, v.Value, v.Value)
	}
	ele := this.l.Front()
	for ele != nil {
		fmt.Printf("list: %+v, %T, \n", ele.Value, ele.Value)
		ele = ele.Next()
	}

}
func (this *LRUCache) len() int {
	return this.l.Len()
}

func (this *LRUCache) isEmpty() bool {
	return this.l.Len() == 0
}

func (this *LRUCache) isFull() bool {
	return this.l.Len() >= this.capacity
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
