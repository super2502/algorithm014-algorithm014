package Week_10

func videoStitching(clips [][]int, T int) int {
	end := 0
	cnt := 0
	lastEnd := 0
	for {
		cnt++
		for _, clip := range clips {
			if clip[0] <= end && clip[1] > lastEnd {
				lastEnd = clip[1]
			}
		}
		if lastEnd >= T {
			return cnt
		}
		if lastEnd == end {
			return -1
		}
		end = lastEnd
	}
}
