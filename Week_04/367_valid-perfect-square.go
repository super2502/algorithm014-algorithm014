package Week_04

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := l + (r-l)/2
		c := mid * mid
		if c == num {
			return true
		} else if c < num {
			l = mid + 1
		} else if c > num {
			r = mid - 1
		}
	}
	return r*r == num
}
