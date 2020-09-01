package Day23

func fib(N int) int {
	if N <= 1 {
		return N
	}
	fir, sec := 0, 1
	var f int
	for i := 2; i <= N; i++ {
		f = fir + sec
		fir = sec
		sec = f
	}

	return f
}
