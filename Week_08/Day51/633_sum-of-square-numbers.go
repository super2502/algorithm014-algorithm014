package Day51

// 为什么双指针比哈希表快这么多？？
func judgeSquareSum(c int) bool {
	i, j := 0, sqrt(c)
	for i <= j {
		pi := i * i
		pj := j * j
		if pi == c-pj {
			return true
		} else if pi < c-pj {
			i++
		} else {
			j--
		}
	}
	return false
}
func sqrt(c int) int {
	l, r := 0, c
	for l <= r {
		mid := l + (r-l)>>1
		prod := mid * mid
		if prod == c {
			return mid
		} else if prod < c {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
func judgeSquareSum1(c int) bool {

	i := 0
	squareMap := make(map[int]bool)
	for {
		sq := i * i
		if sq > c {
			break
		}
		squareMap[sq] = true
		i++
	}
	for sq := range squareMap {
		if squareMap[c-sq] {
			return true
		}
	}
	return false
}
