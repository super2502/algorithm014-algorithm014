package Week_08

func hammingWeight(num uint32) int {
	var cnt int
	for num != 0 {
		cnt++
		num &= num - 1
	}
	return cnt
}
