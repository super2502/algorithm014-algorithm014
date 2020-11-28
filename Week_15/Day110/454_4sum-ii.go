package Day110

func fourSumCount(A []int, B []int, C []int, D []int) int {
	n := len(A)
	map1 := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			map1[A[i]+B[j]]++
		}
	}
	ret := 0
	for k := 0; k < n; k++ {
		for l := 0; l < n; l++ {
			ret += map1[-C[k]-D[l]]
		}
	}
	return ret
}
