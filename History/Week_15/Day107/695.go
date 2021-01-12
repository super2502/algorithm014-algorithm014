package Day107

import (
	"fmt"
)

func findMinMinutes(n int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if n > k {
		return n - k
	}
	if n == k {
		return 0
	}
	if k%2 == 0 {
		if n == k/2 {
			return 1
		} else if n > k/2 {
			return min(n-k/2+1, k-n)
		} else {
			return findMinMinutes(n, k/2) + 1
		}
	} else {
		fmt.Printf("k, n: %v, %v\n", k, n)
		if n == k/2 {
			return 2
		} else if n > k/2 {
			return min(2*n-k+1, min(k-n, n-k/2+2))
		} else {
			return findMinMinutes(n, k/2) + 2
		}
	}

}
