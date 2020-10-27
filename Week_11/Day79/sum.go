package Day79

func sum(a, b int) int {

	for b != 0 {
		//fmt.Printf("%b, %b, %v, %v\n", a, b, a, b)
		//fmt.Printf("%b, %b\n", a &^b , a&b)
		a, b = a^b, (a&b)<<1
	}

	return a
}
