package Day94

var mem [][]int

func findRotateSteps(ring string, key string) int {
	mem = make([][]int, len(ring))
	for i := 0; i < len(ring); i++ {
		mem[i] = make([]int, len(key))
	}

	return helper(0, ring, 0, key)
}

func helper(idx int, ring string, keyIdx int, key string) int {
	if keyIdx == len(key) {
		return 0
	}
	if mem[idx][keyIdx] != 0 {
		return mem[idx][keyIdx]
	}
	b := key[keyIdx]
	left, right := idx, idx
	leftCnt, rightCnt := 0, 0
	for ring[left] != b {
		leftCnt++
		if left > 0 {
			left--
		} else {
			left = len(ring) - 1
		}
	}
	for ring[right] != b {
		rightCnt++
		if right < len(ring)-1 {
			right++
		} else {
			right = 0
		}
	}
	//fmt.Printf("looking for %v, leftCnt: %v, rightCnt: %v\n", string(b), leftCnt, rightCnt)
	var cnt int
	if left == right {
		cnt = min(leftCnt, rightCnt) + helper(left, ring, keyIdx+1, key) + 1
	} else {
		l := helper(left, ring, keyIdx+1, key)
		r := helper(right, ring, keyIdx+1, key)
		cnt = min(l+leftCnt, r+rightCnt) + 1
	}
	mem[idx][keyIdx] = cnt
	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
