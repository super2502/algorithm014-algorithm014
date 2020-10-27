package Day79

import "math"

func divide(dividend int, divisor int) int {
	ret := 0
	a := abs(dividend)
	for a >= abs(divisor) {
		b := abs(divisor)
		cnt2 := 1
		for a >= b+b {
			b = b + b
			cnt2 = cnt2 + cnt2
		}
		ret += cnt2
		a = a - b
	}
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		if ret > math.MaxInt32 {
			return math.MaxInt32
		}
		return ret
	}
	return -ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
