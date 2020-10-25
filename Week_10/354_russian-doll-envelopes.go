package Week_10

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] > envelopes[j][0] {
			return false
		} else {
			return envelopes[i][1] > envelopes[j][1]
		}
	})
	ret := make([]int, 0)
	ret = append(ret, envelopes[0][1])
	for i := 1; i < len(envelopes); i++ {
		target := envelopes[i][1]
		l, r := 0, len(ret)-1
		for l <= r {
			mid := l + (r-l)/2
			if ret[mid] == target {
				l = mid
				break
			} else if ret[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if l == len(ret) {
			ret = append(ret, target)
		} else {
			ret[l] = target
		}
	}

	return len(ret)
}
