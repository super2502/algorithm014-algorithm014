package Day143

func dailyTemperatures(T []int) []int {
	st := make(stack, 0)
	ret := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for !st.IsEmpty() && T[i] > T[st.Peek()] {
			idx := st.Pop()
			ret[idx] = i - idx
		}
		st.Push(i)
	}
	return ret
}

type stack []int

func (st *stack) Len() int {
	return len(*st)
}
func (st *stack) IsEmpty() bool {
	return st.Len() == 0
}
func (st *stack) Push(x int) {
	*st = append(*st, x)
}
func (st *stack) Pop() int {
	x := (*st)[st.Len()-1]
	*st = (*st)[:st.Len()-1]
	return x
}
func (st *stack) Peek() int {
	x := (*st)[st.Len()-1]
	return x
}
