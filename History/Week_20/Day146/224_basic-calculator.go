package Day146

import (
	"fmt"
	"math"
	"strconv"
)

func calculate(s string) int {
	fmt.Printf("calculate %s\n", s)
	var left = math.MinInt64
	s1, s2 := make(stack, 0), make(stack1, 0)
	var numStr string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numStr += string(s[i])
		} else {
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				s1.Push(num)
				fmt.Printf("push %v into s1\n", num)
			}
			switch s[i] {
			case ' ':
			case '+', '-':
				s2.Push(s[i])
				fmt.Printf("push %v into s2\n", string(s[i]))
			case '(':
				s1.Push(left)
				fmt.Printf("push ( into s1\n")
			case ')':
				sum := 0
				for s1.Peek() != left {
					num := s1.Pop()
					//fmt.Printf("s1 pop (%v) \n", num)
					if !s1.IsEmpty() && s1.Peek() != left {
						operator := s2.Pop()
						if operator == '+' {
							sum += num
						} else {
							sum -= num
						}
					} else {
						sum += num
					}
				}
				if !s1.IsEmpty() {
					s1.Pop()
					//fmt.Printf("s1 pop ( \n", )
				}
				s1.Push(sum)
				//fmt.Printf("push sum %v into s1\n", sum)
			}
			numStr = ""
		}
	}
	if numStr != "" {
		num, _ := strconv.Atoi(numStr)
		s1.Push(num)
		//fmt.Printf("push %v into s1\n", num)
	}
	sum := 0
	for !s1.IsEmpty() {
		num := s1.Pop()
		if !s2.IsEmpty() {
			operator := s2.Pop()
			if operator == '+' {
				sum += num
			} else {
				sum -= num
			}
		} else {
			sum += num
		}
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

type stack1 []byte

func (s *stack1) Len() int {
	return len(*s)
}
func (s *stack1) IsEmpty() bool {
	return s.Len() == 0
}
func (s *stack1) Push(x byte) {
	*s = append(*s, x)
}
func (s *stack1) Pop() byte {
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x
}
func (s *stack1) Peek() byte {
	x := (*s)[s.Len()-1]
	return x
}
