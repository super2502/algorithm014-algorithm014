package Week_04

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		c := mid * mid
		if c == x {
			return mid
		} else if c > x {
			right = mid - 1
		} else if c < x {
			left = mid + 1
		}
	}
	return right
}
