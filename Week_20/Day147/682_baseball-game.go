package Day147

import (
	"strconv"
)

func calPoints(ops []string) int {
	st := make(stack, 0)
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "C":
			st.Pop()
		case "D":
			st.Push(st.Peek() * 2)
		case "+":
			st.Push(st.Peek() + st.Peek2())
		default:
			num, _ := strconv.Atoi(ops[i])
			st.Push(num)
		}
		//fmt.Printf("%+v\n", st)
	}
	sum := 0
	for !st.IsEmpty() {
		sum += st.Pop()
	}
	return sum
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
func (s *stack) Peek2() int {
	x := (*s)[s.Len()-2]
	return x
}
