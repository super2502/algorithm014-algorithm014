package Day87

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	l0, r0 := 0, len(intervals)-1
	for l0 <= r0 {
		mid := l0 + (r0-l0)>>1
		if intervals[mid][0] == newInterval[0] {
			r0 = mid
			break
		} else if intervals[mid][0] < newInterval[0] {
			l0 = mid + 1
		} else {
			r0 = mid - 1
		}
	}
	l1, r1 := 0, len(intervals)-1
	for l1 <= r1 {
		mid := l1 + (r1-l1)>>1
		if intervals[mid][1] == newInterval[1] {
			l1 = mid
			break
		} else if intervals[mid][1] < newInterval[1] {
			l1 = mid + 1
		} else {
			r1 = mid - 1
		}
	}
	ret := make([][]int, 0)
	if r0 < 0 && l1 >= len(intervals) {
		return [][]int{newInterval}
	} else if r0 < 0 && l1 < len(intervals) {
		if newInterval[1] < intervals[l1][0] {
			ret = append(ret, newInterval)
			ret = append(ret, intervals[l1:]...)
		} else {
			ret = append(ret, []int{newInterval[0], intervals[l1][1]})
			ret = append(ret, intervals[l1+1:]...)
		}
	} else if r0 >= 0 && l1 >= len(intervals) {
		if newInterval[0] <= intervals[r0][1] {
			ret = append(ret, intervals[:r0]...)
			ret = append(ret, []int{intervals[r0][0], newInterval[1]})
		} else {
			ret = append(ret, intervals[:r0+1]...)
			ret = append(ret, newInterval)
		}
	} else {
		if newInterval[0] > intervals[r0][1] {
			ret = append(ret, intervals[:r0+1]...)
			if newInterval[1] < intervals[l1][0] {
				ret = append(ret, newInterval)
				ret = append(ret, intervals[l1:]...)
			} else {
				ret = append(ret, []int{newInterval[0], intervals[l1][1]})
				ret = append(ret, intervals[l1+1:]...)
			}
		} else {
			ret = append(ret, intervals[:r0]...)
			if newInterval[1] < intervals[l1][0] {
				ret = append(ret, []int{intervals[r0][0], newInterval[1]})
				ret = append(ret, intervals[l1:]...)
			} else {
				ret = append(ret, []int{intervals[r0][0], intervals[l1][1]})
				ret = append(ret, intervals[l1+1:]...)
			}
		}
	}

	return ret
}
