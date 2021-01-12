package Day94

import (
	"math/rand"
)

func kClosest(points [][]int, K int) [][]int {
	dists := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		dists[i] = points[i][0]*points[i][0] + points[i][1]*points[i][1]
	}
	var quickSelect func(l, r int)
	quickSelect = func(l, r int) {
		pivot := rand.Int()%(r-l+1) + l
		dists[r], dists[pivot] = dists[pivot], dists[r]
		points[r], points[pivot] = points[pivot], points[r]
		idx := l
		for i := l; i < r; i++ {
			if dists[i] < dists[r] {
				dists[i], dists[idx] = dists[idx], dists[i]
				points[i], points[idx] = points[idx], points[i]
				idx++
			}
		}
		dists[r], dists[idx] = dists[idx], dists[r]
		points[r], points[idx] = points[idx], points[r]
		if idx == K-1 {
			return
		} else if idx < K-1 {
			quickSelect(idx+1, r)
		} else {
			quickSelect(l, idx-1)
		}
	}
	quickSelect(0, len(points)-1)
	return points[:K]
}
