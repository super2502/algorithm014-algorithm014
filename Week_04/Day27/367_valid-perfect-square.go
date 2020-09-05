package Day27

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := l + (r-l)/2
		y := mid * mid
		if y == num {
			return true
		} else if y > num {
			r = mid - 1
		} else if y < num {
			l = mid + 1
		}
	}
	return false
}
