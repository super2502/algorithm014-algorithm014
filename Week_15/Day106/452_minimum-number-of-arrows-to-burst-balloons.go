package Day106

import "sort"

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return n
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		} else if points[i][0] > points[j][0] {
			return false
		} else {
			return points[i][1] < points[j][1]
		}
	})
	cnt := 1
	_, end := points[0][0], points[0][1]
	for i := 1; i < n; i++ {
		if points[i][0] > end {
			cnt++
			//start = points[i][0]
			end = points[i][1]
		} else {
			//start = points[i][0]
			end = min(end, points[i][1])
		}
	}
	return cnt

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
