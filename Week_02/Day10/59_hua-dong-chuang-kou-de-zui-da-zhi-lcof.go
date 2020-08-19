package Day10

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	l := len(nums)
	ret := make([]int, l-k+1)

	deque := list.New()
	for i := 0; i < k; i++ {
		for deque.Len() > 0 && nums[i] > nums[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(i)
	}
	ret[0] = nums[deque.Front().Value.(int)]

	for i := k; i < l; i++ {
		if deque.Front().Value.(int) == i-k {
			deque.Remove(deque.Front())
		}
		for deque.Len() > 0 && nums[i] > nums[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(i)
		ret[i-k+1] = nums[deque.Front().Value.(int)]
	}
	return ret

}
