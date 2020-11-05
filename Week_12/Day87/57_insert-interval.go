package Day87

func insert1(array [][]int, rang []int) [][]int {
	i, n := 0, len(array)
	ret := make([][]int, 0)
	mid := make([]int, 2)

	for i < n && array[i][1] < rang[0] {
		i++
	}
	ret = append(ret, array[:i]...)
	if i == n {
		ret = append(ret, rang)
		return ret
	}
	mid[0] = min(array[i][0], rang[0])
	for i < n && array[i][0] < rang[1] {
		i++
	}
	if i == 0 {
		ret = append([][]int{rang}, array[i:]...)
		return ret
	}
	mid[1] = max(array[i-1][1], rang[1])
	ret = append(ret, mid)
	ret = append(ret, array[i:]...)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
