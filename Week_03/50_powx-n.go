package Week_03

func myPow(x float64, n int) float64 {

	if n < 0 {
		x = 1 / x
		n = -1 * n
	}
	return halfPow(x, n)
}

func halfPow(x float64, n int) float64 {

	a := 1.0
	if n%2 == 1 {
		a = x
	}
	if n/2 == 0 {
		return a
	}
	y := halfPow(x, n/2)

	return a * y * y
}
