package Week_01

type MyCircularDeque struct {
	arr      []int
	capacity int
	head     int
	rear     int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {

	return MyCircularDeque{
		arr:      make([]int, k+1),
		head:     0,
		rear:     0,
		capacity: k,
	}

}

func (this *MyCircularDeque) next(cur int) int {
	return (cur + 1) % (this.capacity + 1)
}
func (this *MyCircularDeque) last(cur int) int {
	return (cur + this.capacity) % (this.capacity + 1)
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = this.last(this.head)
	this.arr[this.head] = value

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.arr[this.rear] = value
	this.rear = this.next(this.rear)

	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = this.next(this.head)

	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {

	if this.IsEmpty() {
		return false
	}
	this.rear = this.last(this.rear)

	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {

	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {

	if this.IsEmpty() {
		return -1
	}

	last := this.last(this.rear)
	return this.arr[last]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {

	return this.head == this.rear
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {

	return (this.rear+1)%(this.capacity+1) == this.head

}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
