package Day83

import (
	"math/rand"
)

type RandomizedCollection struct {
	list []int
	m    map[int]map[int]struct{}
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		list: make([]int, 0),
		m:    make(map[int]map[int]struct{}),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	notExists := len(this.m[val]) <= 0
	this.list = append(this.list, val)
	if len(this.m[val]) == 0 {
		this.m[val] = make(map[int]struct{})
	}
	this.m[val][len(this.list)-1] = struct{}{}
	return notExists
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if len(this.m[val]) <= 0 {
		return false
	}
	var valIdx int
	for idx := range this.m[val] {
		valIdx = idx
		break
	}
	lastIdx := len(this.list) - 1
	lastVal := this.list[lastIdx]
	this.list[valIdx], this.list[lastIdx] = this.list[lastIdx], this.list[valIdx]
	delete(this.m[lastVal], lastIdx)
	this.m[lastVal][valIdx] = struct{}{}
	delete(this.m[val], valIdx)
	this.list = this.list[:len(this.list)-1]
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	idx := rand.Int() % len(this.list)
	return this.list[idx]
}
