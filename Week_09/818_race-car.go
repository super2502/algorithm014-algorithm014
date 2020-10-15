package Week_09

import (
	"fmt"
)

func racecar(target int) int {

	meters := make([]int, target + 1)
	speed := 1
	wp := 0
	nextWp := 0
	for i := 0; i <= target; i++ {
		if i == nextWp {
			wp = nextWp
			nextWp = wp + speed
			speed *= 2
			if nextWp <= target {
				meters[nextWp] = meters[wp] + 1
			}
		} else {
			fromLeft := meters[wp] + 2 + meters[i-wp]
			fromRight := meters[wp] + 1 + 1 + meters[nextWp - i]
			meters[i] = min(fromLeft, fromRight)
		}
	}
	fmt.Printf("%+v, %v\n", meters, nextWp)

	return meters[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}