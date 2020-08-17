package Day3

import (
	"container/list"
)

func maxSlidingWindow1(nums []int, k int) []int {

	ret := make([]int, 0)
	deque := list.New()
	//fmt.Printf("deque len (%v)\n", deque.Len())
	deque.PushBack(nums[0])
	//fmt.Printf("deque len (%v)\n", deque.Len())
	l := len(nums)

	for i := 1; i < k; i++ {
		for deque.Len() > 0 {
			right := deque.Back()
			//fmt.Printf("deque len (%v) right %v, %T\n", deque.Len(),right.Value, right.Value)
			if nums[i] > right.Value.(int) {
				deque.Remove(right)
			} else {
				break
			}
		}
		deque.PushBack(nums[i])
		//fmt.Printf("deque len (%v)(%v)\n", deque.Len(), deque.Front().Value)
	}
	ret = append(ret, deque.Front().Value.(int))
	for i := k; i < l; i++ {
		//fmt.Printf("i:(%v), len(%v)\n", i, deque.Len())
		if deque.Front().Value.(int) == nums[i-k] {
			_ = deque.Remove(deque.Front())
			//fmt.Printf("(%v) pop (%v)\n", i, pop)
		}
		for deque.Len() > 0 {
			right := deque.Back()
			if nums[i] > right.Value.(int) {
				deque.Remove(right)
			} else {
				break
			}
		}
		deque.PushBack(nums[i])
		ret = append(ret, deque.Front().Value.(int))
	}

	return ret
}
