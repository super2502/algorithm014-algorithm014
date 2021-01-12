package Day117

func lengthOfLIS(nums []int) int {
	ret := make([]int, 0)
	for _, num := range nums {
		l, r := 0, len(ret)-1
		for l <= r {
			mid := l + (r-l)>>1
			if ret[mid] == num {
				l = mid
				break
			} else if ret[mid] < num {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if l == len(ret) {
			ret = append(ret, num)
		} else {
			ret[l] = num
		}
	}
	return len(ret)
}
