package Day3

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0)
	l := len(nums)
	if l == 0 || k == 0 {
		return ret
	}
	de := make(deque, 0)
	//fmt.Printf("%+v, %v\n", nums, k)
	for i := 0; i < k; i++ {
		for !de.IsEmpty() && nums[i] > nums[de.Back()] {
			_ = de.PopBack()
		}
		de.PushBack(i)

	}
	ret = append(ret, nums[de.Front()])
	for i := k; i < l; i++ {

		//fmt.Printf("(%v) (%v)\n", i, de)
		if !de.IsEmpty() && de.Front() == (i-k) {
			_ = de.PopFront()
		}
		for !de.IsEmpty() && nums[i] > nums[de.Back()] {
			_ = de.PopBack()
		}
		de.PushBack(i)
		ret = append(ret, nums[de.Front()])
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 自己仿制一个简易的deque
type deque []int

func (d *deque) Len() int {
	return len(*d)
}

func (d *deque) IsEmpty() bool {

	return len(*d) == 0
}
func (d *deque) Front() int {

	return (*d)[0]
}

func (d *deque) Back() int {

	return (*d)[d.Len()-1]
}

func (d *deque) PushBack(x int) {

	*d = append(*d, x)
}

func (d *deque) PopBack() int {

	x := (*d)[d.Len()-1]
	*d = (*d)[:d.Len()-1]
	return x
}

func (d *deque) PopFront() int {
	x := (*d)[0]
	*d = (*d)[1:d.Len()]
	return x
}

func (d *deque) String() string {
	return fmt.Sprintf("%+v", *d)
}
