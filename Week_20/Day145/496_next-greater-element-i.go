package Day145

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	st := make(stack, 0)
	m2 := make(map[int]int)
	for _, num2 := range nums2 {
		for !st.IsEmpty() && st.Peek() < num2 {
			num := st.Pop()
			m2[num] = num2
		}
		st.Push(num2)
	}
	for !st.IsEmpty() {
		num := st.Pop()
		m2[num] = -1
	}
	ret := make([]int, len(nums1))
	for idx, num1 := range nums1 {
		ret[idx] = m2[num1]
	}
	return ret
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}
func (s *stack) Push(x int) {
	*s = append(*s, x)
}
func (s *stack) Pop() int {
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x
}
func (s *stack) Peek() int {
	x := (*s)[s.Len()-1]
	return x
}
