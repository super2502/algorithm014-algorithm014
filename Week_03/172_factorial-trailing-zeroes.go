package Week_03

import "fmt"

func trailingZeroes(n int) int {

	for i := 0; i < 130; i++ {
		fmt.Printf("%v, %v\n", i, a(i))
	}

	return a(n)
}

func a(n int) int {
	sum := 0
	for n > 0 {
		n = n / 5
		sum += n
	}

	return sum
}
