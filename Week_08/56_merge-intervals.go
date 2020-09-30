package Week_08

import "sort"

func merge56(intervals [][]int) [][]int {
	ret := make([][]int, 0)
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= last[1] {
			if last[1] < intervals[i][1] {
				last[1] = intervals[i][1]
			}
		} else {
			ret = append(ret, last)
			last = intervals[i]
		}
	}
	ret = append(ret, last)
	return ret
}
