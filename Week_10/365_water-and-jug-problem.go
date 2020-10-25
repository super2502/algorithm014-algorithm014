package Week_10

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if z > x+y {
		return false
	}
	if z%gcd(x, y) == 0 {
		return true
	}
	return false
}

func gcd(a, b int) int {
	if a > b {
		a, b = b, a
	}
	if a == 0 {
		return b
	}
	return gcd(a, b-a)
}
