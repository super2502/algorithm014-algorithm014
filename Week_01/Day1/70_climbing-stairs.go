package Day1

import (
	"fmt"
	"testing"
)

var _ = fmt.Errorf("")

func BenchmarkAaa(b *testing.B) {

	for j := 1; j < b.N; j++ {
		cache := make(map[int]int)
		A(10000, cache)
	}
}

func BenchmarkB(b *testing.B) {

	for j := 1; j < b.N; j++ {
		B(10000)
	}
}

func A(n int, cache map[int]int) int {
	if n == 1 {
		cache[1] = 1
		return 1
	}
	if n == 2 {
		cache[1] = 1
		cache[2] = 2
		return 2
	}
	a, ok := cache[n-1]
	if !ok {
		a = A(n-1, cache)
	}
	b, ok := cache[n-2]
	if !ok {
		b = A(n-2, cache)
	}
	c := a + b
	cache[n] = c
	return c
}

func B(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 1, 2, 0
	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
		//fmt.Printf("a,b,c %v %v %v\n", a, b, c)
	}
	return c
}

func TestAB(t *testing.T) {

	cache := make(map[int]int)
	for i := 10; i < 20; i++ {
		a := A(i, cache)
		b := B(i)
		t.Logf("a, b: %v ,%v", a, b)
	}
}
